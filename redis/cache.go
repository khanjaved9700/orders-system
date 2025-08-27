package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
}

type redisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewCache(addr string) Cache {
	rdb := redis.NewClient(&redis.Options{Addr: addr})
	return &redisCache{client: rdb, ctx: context.Background()}
}

func (c *redisCache) Set(key, value string) error {
	return c.client.Set(c.ctx, key, value, 0).Err()
}

func (c *redisCache) Get(key string) (string, error) {
	return c.client.Get(c.ctx, key).Result()
}
