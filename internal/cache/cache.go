package cache

import (
    "context"
    "time"
    "github.com/redis/go-redis/v9"
)

type Cache struct {
    rdb *redis.Client
}

func New(addr string) *Cache {
    return &Cache{rdb: redis.NewClient(&redis.Options{Addr: addr})}
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
    return c.rdb.Get(ctx, key).Result()
}

func (c *Cache) Set(ctx context.Context, key string, val []byte, ttl time.Duration) error {
    return c.rdb.Set(ctx, key, val, ttl).Err()
}

func (c *Cache) TTL(ctx context.Context, key string) (time.Duration, error) {
    return c.rdb.TTL(ctx, key).Result()
}
