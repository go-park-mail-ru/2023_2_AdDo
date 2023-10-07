package init_redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitRedis(ctx context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "service-db-redis:6379", // Адрес и порт вашего Redis-сервера
		Password: "",                      // Пароль Redis-сервера (если требуется)
		DB:       0,                       // Номер используемой Redis-базы данных
	})
	fmt.Println("Redis database client created!")

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("Redis database successfully connected!")

	return rdb, nil
}
