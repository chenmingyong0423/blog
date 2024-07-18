package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient 初始化 Redis 客户端
func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// GetConsecutiveDays 计算连续签到天数
func GetConsecutiveDays(ctx context.Context, rdb *redis.Client, userID int, year int, dayOfYear int) (int, error) {
	key := fmt.Sprintf("user:%d:%d", year, userID)
	segmentSize := 63
	consecutiveDays := 0
	bitOps := make([]any, 0)

	for i := 0; i < dayOfYear; i += segmentSize {
		size := segmentSize
		if i+segmentSize > dayOfYear {
			size = dayOfYear - i
		}

		bitOps = append(bitOps, "GET", fmt.Sprintf("u%d", size), fmt.Sprintf("#%d", i))
	}

	values, err := rdb.BitField(ctx, key, bitOps...).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get bitfield: %w", err)
	}

	for idx, value := range values {
		if value != 0 {
			size := segmentSize
			if (idx+1)*segmentSize > dayOfYear {
				size = dayOfYear % segmentSize
			}
			for j := 0; j < size; j++ {
				if (value & (1 << (size - 1 - j))) != 0 {
					consecutiveDays++
				}
			}
		}
	}
	return consecutiveDays, nil
}

func main() {
	rdb := RedisClient()
	if rdb == nil {
		log.Fatal("redis client is nil")
	}
	now := time.Now()
	// 获取当前的年份
	year := now.Year()
	// 获取当前日期是今年的第几天
	dayOfYear := now.YearDay()
	// 假设用户 ID 为 1
	userID := 1

	consecutiveDays, err := GetConsecutiveDays(context.Background(), rdb, userID, year, dayOfYear)
	if err != nil {
		log.Fatalf("failed to get consecutive days: %v", err)
	}

	fmt.Printf("%d 年累计签到的天数: %d\n", year, consecutiveDays)
}
