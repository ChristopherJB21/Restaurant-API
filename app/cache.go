package app

import (
	"restaurant/model/web"

	"github.com/klauspost/compress/zstd"
	"github.com/redis/go-redis/v9"
)

func NewCustomCache(redis *redis.Client, zstdWriter *zstd.Encoder, zstdReader *zstd.Decoder) *web.CustomCache {
	return &web.CustomCache{
		Redis:      redis,
		ZstdWriter: zstdWriter,
		ZstdReader: zstdReader,
	}
}
