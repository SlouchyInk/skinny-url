package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(host string) *RedisCache {
	client := redis.NewClient(&redis.Options{Addr: host})
	return &RedisCache{client: client}
}

func (r *RedisCache) Set(key, value string) error {
	return r.client.Set(context.Background(), key, value, time.Hour*24*7).Err()
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}
