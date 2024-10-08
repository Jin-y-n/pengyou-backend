package test

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestRedis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 10,
	})
	// info := rdb.Info(context.Background())

	// fmt.Println(info)
	// t.Log(info)

	t.Log(rdb.Info(context.Background(), "used", "_", "memory"))
}
