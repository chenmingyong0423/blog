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
	oldValue, err := rdb.SetBit(context.Background(), "user:2024:1", 0, 1).Result()
	if err != nil {
		panic(err)
	}
	if oldValue == 1 {
		fmt.Println("重复签到")
	} else {
		fmt.Println(oldValue) // 0，表示这个位（`bit`）被设置新值之前的值。
	}
}
