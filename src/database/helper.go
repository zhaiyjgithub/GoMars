package database

import (
	"GoMars/src/conf"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	redisEngineOnce sync.Once
	redisEngine *redis.Client
)

func SetupRedisEngine() *redis.Client {
	redisEngineOnce.Do(func() {
		c := conf.RedisConf
		addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
		redisEngine = redis.NewClient(&redis.Options{
			Addr: addr,
			Password: "",
			DB: 0,
		})

		_, err := redisEngine.Ping(context.Background()).Result()
		if err != nil {
			fmt.Printf("Connect failed with Redis: %v \n", err)
		}else {
			fmt.Printf("Connect to redis: %s\r\n", addr)
		}
	})

	return  redisEngine
}
