package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type CacheInterface interface {
	Get(key string) ([]byte, error)
	Set(key string, state []byte, expire ...int) error
	Del(key string) error
	Expire(key string) error
}

type CacheHandler struct {
	client *redis.Client
}

func NewCacheHandler() (CacheInterface, error) {
	redisHostPort := "127.0.0.1:6379"
	return &CacheHandler{
		client: redis.NewClient(&redis.Options{
			Addr:         redisHostPort,
			MinIdleConns: 5,
			MaxRetries:   3,
			Password:     "",
			DB:           0,
		}),
	}, nil
}

func (cs CacheHandler) Get(key string) ([]byte, error) {
	val, err := cs.client.Get(ctx, key).Bytes()
	if err != nil {
		return []byte("UNKNOWN"), err
	}
	return val, nil
}

func (cs CacheHandler) Set(key string, state []byte, expired ...int) error {
	expire := 0 * time.Second
	if len(expired) > 0 {
		expire = time.Duration(expired[0])
	}
	err := cs.client.Set(ctx, key, state, expire*time.Second).Err()
	return err
}

func (cs CacheHandler) Del(key string) error {
	return cs.client.Del(ctx, key).Err()
}

func (cs CacheHandler) Expire(key string) error {
	return cs.client.Expire(ctx, key, 10*time.Second).Err()
}
