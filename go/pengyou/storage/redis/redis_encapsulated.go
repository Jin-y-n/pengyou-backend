package rds

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"pengyou/global/config"
	"pengyou/utils/log"
	"strconv"
	"strings"
	"time"
)

// ZAdd this file encapsulates redis client functions
func ZAdd(context context.Context, key string, member ...redis.Z) {
	if RedisClient != nil {
		RedisClient.ZAdd(context, key, member...)
	} else if RedisClusterClient != nil {
		RedisClusterClient.ZAdd(context, key, member...)
	}
}

func ZRangeByScore(context context.Context, key string, min string, max string) ([]string, error) {
	var err error
	var result []string

	if RedisClient != nil {
		result, err = RedisClient.ZRangeByScore(
			context,
			key,
			&redis.ZRangeBy{
				Min: min,
				Max: max,
			}).Result()
	} else if RedisClusterClient != nil {
		result, err = RedisClusterClient.ZRangeByScore(context, key,
			&redis.ZRangeBy{Min: min, Max: max}).Result()
	}

	return result, err
}

func RedisPublish(context context.Context, channel string, message string) error {
	var err error
	if RedisClient != nil {
		err = RedisClient.Publish(context, channel, message).Err()
		log.Logger.Info("publish message to redis: " + message)
		return err
	} else if RedisClusterClient != nil {
		err = RedisClusterClient.Publish(context, channel, message).Err()
		log.Logger.Info("publish message to redis cluster: " + message)
		return err
	}

	log.Logger.Error("redis not init")

	return err
}

func RedisPublishObj(context context.Context, channel string, obj interface{}) error {
	var err error

	marshalMessage, err := json.Marshal(obj)
	if err != nil {
		log.Logger.Error("marshal object failed", zap.Error(err))
		return err
	}

	if RedisClient != nil {
		err = RedisClient.Publish(context, channel, marshalMessage).Err()
		log.Logger.Info("publish marshalMessage to redis: " + string(marshalMessage))
		return err
	} else if RedisClusterClient != nil {
		err = RedisClusterClient.Publish(context, channel, marshalMessage).Err()
		log.Logger.Info("publish marshalMessage to redis cluster: " + string(marshalMessage))
		return err
	}

	log.Logger.Error("redis not init")

	return err
}

func NativeSubscribe(contecct context.Context, channel string) *redis.PubSub {
	if RedisClient != nil {
		return RedisClient.Subscribe(contecct, channel)
	} else if RedisClusterClient != nil {
		return RedisClusterClient.Subscribe(contecct, channel)
	}

	return nil
}

func RedisSubscribe(context context.Context, channel string, callback func(message string)) error {
	var err error

	if RedisClient != nil {
		sub := RedisClient.Subscribe(context, channel)
		_, err = sub.Receive(context)
		if err != nil {
			return err
		}

		msg, err := sub.ReceiveMessage(context)
		// cmd :=
		// RedisClient.LPop(context, channel)

		// if cmd.Err() != nil {
		// 	log.Logger.Error("receive message from redis error: " + err.Error())
		// }

		if err != nil {
			log.Logger.Error("receive message from redis error: " + err.Error())
			return err
		}

		callback(msg.Payload)
	} else if RedisClusterClient != nil {
		sub := RedisClusterClient.Subscribe(context, channel)
		defer sub.Close()

		_, err = sub.Receive(context)
		if err != nil {
			return err
		}

		msg, err := sub.ReceiveMessage(context)
		// cmd :=
		// RedisClient.LPop(context, channel)

		// if cmd.Err() != nil {
		// 	log.Logger.Error("receive message from redis error: " + err.Error())
		// }

		if err != nil {
			return err
		}

		callback(msg.Payload)
	}

	return err
}

func GenerateName(userId uint) string {
	return config.Cfg.App.PublishKey + ":to" + fmt.Sprint(userId)
}

func RedisInfo(args ...string) *redis.StringCmd {
	if RedisClusterClient != nil {
		return RedisClusterClient.Info(context.Background(), args...)
	} else {
		return RedisClient.Info(context.Background(), args...)
	}
}

func GetRedisMemoryUsed() (int64, error) {
	infos, err := RedisInfo("memory", "used").Result()

	if err != nil {
		return 0, err
	}

	for _, line := range strings.Split(infos, "\n") {
		if strings.HasPrefix(line, "used_memory:") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				return strconv.ParseInt(parts[1], 10, 64) // Return the memory usage as a string
			}
		}
	}

	return 0, errors.New("memory usage not found")
}

func Get(context context.Context, key string) *redis.StringCmd {
	if RedisClient != nil {
		return RedisClient.Get(context, key)
	} else if RedisClusterClient != nil {
		return RedisClusterClient.Get(context, key)
	}

	log.Logger.Error("redis not init")

	return nil
}

func Set(context context.Context, key string, value string) *redis.StatusCmd {
	if RedisClient != nil {
		return RedisClient.Set(context, key, value, 0)
	} else if RedisClusterClient != nil {
		return RedisClusterClient.Set(context, key, value, 0)
	}

	return nil
}

func SetWithExpire(context context.Context, key string, value string, expire time.Duration) *redis.StatusCmd {
	if RedisClient != nil {
		return RedisClient.Set(context, key, value, expire)
	} else if RedisClusterClient != nil {
		return RedisClusterClient.Set(context, key, value, expire)
	}

	return nil
}

func SetObj(context context.Context, key string, value interface{}) *redis.StatusCmd {
	return SetObjWithExpire(context, key, value, 0)
}

func SetObjWithExpire(context context.Context, key string, value interface{}, expire time.Duration) *redis.StatusCmd {
	bytes, err := json.Marshal(value)
	res := &redis.StatusCmd{}

	if err != nil {
		res.SetErr(err)
		return res
	}

	if RedisClient != nil {
		return RedisClient.Set(context, key, bytes, expire)
	} else if RedisClusterClient != nil {
		return RedisClusterClient.Set(context, key, bytes, expire)
	}

	res.SetErr(errors.New("redis not init"))
	return res
}

// this function is used to find all keys with the given prefix
func ScanKeysWithPrefix(prefix string) ([]string, error) {
	var keys []string
	cursor := uint64(0)
	var result *redis.ScanCmd
	for {
		result = RedisClient.Scan(context.Background(), cursor, prefix+"*", 10)
		err := result.Err()
		if err != nil {
			return nil, err
		}
		keysPart, _ := result.Val()
		keys = append(keys, keysPart...)
		if cursor == 0 {
			break
		}
	}
	return keys, nil
}

func Del(key string) {
	if RedisClient != nil {
		RedisClient.Del(context.Background(), key)
	}
}
