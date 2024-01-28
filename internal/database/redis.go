package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nebojsaj1726/movie-base/config"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() (*RedisClient, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return &RedisClient{Client: rdb}, nil
}

func (r *RedisClient) SetCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	resultJSON, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error marshaling value: %v", err)
	}

	cmd := r.Client.Set(ctx, key, resultJSON, expiration)
	return cmd.Err()
}

func (r *RedisClient) GetCache(ctx context.Context, key string, resultModel interface{}) (interface{}, error) {
	cmd := r.Client.Get(ctx, key)
	result, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(result), resultModel); err != nil {
		return nil, err
	}

	return resultModel, nil
}

func (r *RedisClient) DeleteCache(ctx context.Context, key string) error {
	cmd := r.Client.Del(ctx, key)
	return cmd.Err()
}
