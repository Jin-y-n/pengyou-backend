package test

import (
	"pengyou/global/config"
	rds "pengyou/storage/redis"
	"testing"
)

func TestRedis(t *testing.T) {
	rds.InitRedis(&config.Redis{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 10,
	})
	// info := rdb.Info(context.Background())

	// fmt.Println(info)
	// t.log(info)

	//t.log(rds.Info(context.Background(), "used", "_", "memory"))
}
