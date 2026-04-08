package web

import (
	"github.com/klauspost/compress/zstd"
	"github.com/redis/go-redis/v9"
)

type CustomCache struct {
	Redis      *redis.Client
	ZstdWriter *zstd.Encoder
	ZstdReader *zstd.Decoder
}
