package cache

import (
	"context"

	"github.com/hdurham99/skinny-url/internal/storage/db"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(host string) *RedisCache {
	client := redis.NewClient(&redis.Options{Addr: host})
	return &RedisCache{client: client}
}

func (r *RedisCache) Set(short_code, long_url, user_id string) error {
	data := map[string]string{
		"long_url":    long_url,
		"user_id":     user_id,
		"click_count": "null",
	}
	return r.client.HSet(context.Background(), short_code, data).Err()
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.client.HGet(context.Background(), key, "long_url").Result()
}

func (r *RedisCache) IncrementClickCount(short_code string) error {
	return r.client.HIncrBy(context.Background(), short_code, "click_count", 1).Err()
}

func (r *RedisCache) FlushClickCountsToDB(c *db.CassandraDB) {
	keys, _ := r.client.Keys(context.Background(), "*").Result()
	for _, key := range keys {
		count, _ := r.client.Get(context.Background(), key).Int()
		c.SaveClickCount(key, count)
		r.client.Del(context.Background(), key)
	}
}
