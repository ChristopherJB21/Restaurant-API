package helper

import (
	"context"
	"encoding/json"
	"restaurant/model/web"
	"time"
)

func SetCache(ctx context.Context, key string, value interface{}, expiration time.Duration, customCache *web.CustomCache) {
	cacheJSON, err := json.Marshal(value)
	PanicIfError(err)

	compressedCache := customCache.ZstdWriter.EncodeAll(cacheJSON, nil)

	err = customCache.Redis.Set(ctx, key, compressedCache, expiration).Err()
	PanicIfError(err)
}

func GetCache(ctx context.Context, key string, result interface{}, customCache *web.CustomCache) bool {
	getCache := customCache.Redis.Get(ctx, key)

	if getCache.Err() != nil {
		return false
	}

	compressedCache, err := getCache.Bytes()
	PanicIfError(err)

	cacheJSON, err := customCache.ZstdReader.DecodeAll(compressedCache, nil)
	PanicIfError(err)

	err = json.Unmarshal(cacheJSON, result)
	PanicIfError(err)

	return true
}
