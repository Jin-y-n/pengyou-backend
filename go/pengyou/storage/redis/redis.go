package rds

import (
	"context"

	"pengyou/global/config"
	"pengyou/utils/log"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient        *redis.Client
	RedisClusterClient *redis.ClusterClient
)

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
