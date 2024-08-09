package app

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func NewRedis() *redis.Client {
	host := viper.GetString("redis.host")

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		DB:       0,
	})

	return client
}
