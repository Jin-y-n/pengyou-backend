package storage

import (
	"context"
	"errors"
	"fmt"
	"pengyou/global/config"
	"pengyou/utils/log"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient        *redis.Client
	RedisClusterClient *redis.ClusterClient
)

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
		log.Debug("publish message to redis: " + message)
		return err
	} else if RedisClusterClient != nil {
		err = RedisClusterClient.Publish(context, channel, message).Err()
		log.Debug("publish message to redis cluster: " + message)
		return err
	}

	log.Error("redis not init")

	return err
}

func RedisSubsrcibe(context context.Context, channel string, callback func(message string)) error {
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
		// 	log.Error("receive message from redis error: " + err.Error())
		// }

		if err != nil {
			log.Error("receive message from redis error: " + err.Error())
			return err
		}

		callback(msg.Payload)
	}

	return err
}

func InitRedis(cfg *config.Redis) {
	if cfg.UseCluster {
		NewRedisClusterClient(cfg)

		status := RedisClusterClient.Ping(context.Background())
		if status.Err() != nil {
			panic(status.Err())
		}
	} else {
		NewRedisClient(cfg)

		status := RedisClient.Ping(context.Background())
		if status.Err() != nil {
			panic(status.Err())
		}
	}

	var redisType string
	if cfg.UseCluster {
		redisType = "cluster"
	} else {
		redisType = "standalone"
	}

	log.Info("redis init success with " +
		redisType +
		" -> " +
		cfg.Addr)
}

func NewRedisClient(cfg *config.Redis) {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	RedisClient.Ping(context.Background())
}

func NewRedisClusterClient(cfg *config.Redis) {
	RedisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.ClusterAddrs,
		Username: cfg.Username,
		Password: cfg.Password,
		PoolSize: cfg.PoolSize,
	})

	RedisClusterClient.Ping(context.Background())
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
