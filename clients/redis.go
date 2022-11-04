package clients

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

func (rc *RedisClient) Configure() {
	rc.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rc.client.Ping(rc.client.Context()).Result()

	if err != nil {
		panic(err)
	}
}

func (rc *RedisClient) GetValue(key string) string {
	rc.Configure()
	defer rc.client.Close()
	ctx := context.Background()
	value, _ := rc.client.Get(ctx, key).Result()
	return value
}

func (rc *RedisClient) SetValue(key string, val interface{}, expiry time.Duration) string {
	rc.Configure()
	defer rc.client.Close()
	ctx := context.Background()
	value, _ := rc.client.Set(ctx, key, val, expiry).Result()
	return value
}
