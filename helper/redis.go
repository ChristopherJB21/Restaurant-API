package helper

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetCache(ctx context.Context, key string, value interface{}, expiration time.Duration, redis *redis.Client) {
	cacheJSON, err := json.Marshal(value)
	PanicIfError(err)

	err = redis.Set(ctx, key, cacheJSON, expiration).Err()
	PanicIfError(err)
}

func GetCache(ctx context.Context, key string, result interface{}, redis *redis.Client) bool {
	cacheJSON, err := redis.Get(ctx, key).Result()

	if err != nil {
		return false
	}

	err = json.Unmarshal([]byte(cacheJSON), result)
	PanicIfError(err)

	return true
}
