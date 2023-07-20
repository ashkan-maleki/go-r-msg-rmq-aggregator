package redis

import (
	"context"
	redis "github.com/redis/go-redis/v9"
)

type Redis struct {
	redis  *redis.Client
	prefix string
}

func NewRedis(rdb *redis.Client, prefix string) chepun.StateManager {
	if prefix != "" && prefix[len(prefix)-1] != ':' {
		prefix += ":"
	}
	return &Redis{redis: rdb, prefix: prefix}
}

func (r *Redis) Append(ctx context.Context, correlationId string, correlationConfig []byte, queueName string, message []byte) error {
	pipeline := r.redis.Pipeline()
	pipeline.HSet(ctx, r.prefix+correlationId, queueName, message)
	pipeline.HSetNX(ctx, r.prefix+correlationId, chepun.CorrelationConfigKey, correlationConfig)
	_, err := pipeline.Exec(ctx)
	return err
}

func (r *Redis) Delete(ctx context.Context, correlationId string) error {
	return r.redis.Del(ctx, r.prefix+correlationId).Err()
}

func (r *Redis) All(ctx context.Context, correlationId string) (map[string][]byte, error) {
	items, err := r.redis.HGetAll(ctx, r.prefix+correlationId).Result()
	if err != nil {
		return nil, nil
	}
	response := make(map[string][]byte, len(items))
	for key, value := range items {
		response[key] = []byte(value)
	}
	return response, nil
}
