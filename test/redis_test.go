package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

func BenchmarkRedisConnect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rdb := redis.NewClient(&redis.Options{
			Addr:     "redis.suqf.top:6379",
			Password: "root",
			DB:       0,
		})
		_ = rdb.Close()
	}
}

func BenchmarkSet(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis.suqf.top:6379",
		Password: "root",
		DB:       0,
	})
	defer func(rdb *redis.Client) {
		_ = rdb.Close()
	}(rdb)

	for i := 0; i < b.N; i++ {
		cmd := rdb.Set(context.Background(), fmt.Sprintf("ben-%d", i), i, -1)
		b.Log(cmd.String())
	}
}
