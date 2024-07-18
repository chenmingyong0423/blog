package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

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
	now := time.Now()
	// 获取当前的年份
	year := now.Year()
	// 假设用户 ID 为 1
	userID := 1
	// 获取当前月的天数
	days := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location()).Add(-24 * time.Hour).Day()
	// 获取本月初是今年的第几天
	offset := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).YearDay()
	signOfMonth, err := GetSignOfMonth(context.Background(), rdb, userID, year, days, offset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signOfMonth)
}

func GetSignOfMonth(ctx context.Context, rdb *redis.Client, userID, year, days, offset int) ([]bool, error) {
	typ := fmt.Sprintf("u%d", days)
	key := fmt.Sprintf("user:%d:%d", year, userID)

	s, err := rdb.BitField(ctx, key, "GET", typ, offset).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get bitfield: %w", err)
	}

	if len(s) != 0 {
		signInBits := s[0]
		signInSlice := make([]bool, days)
		for i := 0; i < days; i++ {
			signInSlice[i] = (signInBits & (1 << (days - 1 - i))) != 0
		}
		return signInSlice, nil
	} else {
		return nil, errors.New("no result returned from BITFIELD command")
	}
}
