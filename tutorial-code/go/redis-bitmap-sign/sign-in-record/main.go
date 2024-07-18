package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main() {
	rdb := RedisClient()
	if rdb == nil {
		panic("redis client is nil")
	}
	value, err := rdb.GetBit(context.Background(), "user:2024:1", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(value) // 1
}
