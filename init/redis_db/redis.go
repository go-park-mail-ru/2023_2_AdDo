package init_redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	fmt.Println("Redis database client created!")

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("Redis database successfully connected!")

	return rdb, nil
}
