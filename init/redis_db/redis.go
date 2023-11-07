package init_redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "service-db-redis:6379",
		Password: "",
		DB:       0,
	})
	fmt.Println("Redis db client created!")

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	fmt.Println("Redis db successfully connected!")

	return rdb, nil
}
