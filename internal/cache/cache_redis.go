package cache

import (
	"context"
	"encoding/json"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

var _ Cache = &Redis{}

func (c *Redis) Get(ctx context.Context, key string, val any) error {
	resp := c.client.Get(ctx, key)
	if resp.Err() != nil {
		return resp.Err()
	}

	return json.Unmarshal([]byte(resp.Val()), val)
}

func (c *Redis) Set(ctx context.Context, key string, val any, ttl time.Duration) error {
	req, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, req, ttl).Err()
}

func (c *Redis) Has(ctx context.Context, key string) (bool, error) {
	resp := c.client.Exists(ctx, key)
	if resp.Err() != nil {
		return false, resp.Err()
	}

	return resp.Val() > 0, nil
}

func (c *Redis) Keys(ctx context.Context, pattern string) ([]string, error) {
	resp := c.client.Keys(ctx, pattern)
	if resp.Err() != nil {
		return []string{}, resp.Err()
	}

	return resp.Val(), nil
}
