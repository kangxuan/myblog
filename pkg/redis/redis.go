package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"myblog/settings"
)

var client *redis.Client

func SetUp() {
	redisConf := settings.ServerConf.RedisConfig
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		DB:   redisConf.Db,
	})

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Errorf("redis Ping field: %s", err))
	}
}
