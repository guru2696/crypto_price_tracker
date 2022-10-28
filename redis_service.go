package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

func GetRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(redisClient.Context()).Result()

	if err != nil {
		panic(err)
	}
	return redisClient
}

func CheckRedisForPrice(client *redis.Client) map[string]map[string]string {
	ctx := context.Background()
	valUSD, _ := client.Get(ctx, "bitcoin_price_USD").Result()
	valEUR, _ := client.Get(ctx, "bitcoin_price_EUR").Result()

	if valUSD != "" {
		log.Default().Println("Fetching Redis....")
		formattedData := map[string]map[string]string{
			"bitcoin": {
				"EUR": valEUR,
				"USD": valUSD,
			},
		}
		return formattedData
	}
	return nil
}
