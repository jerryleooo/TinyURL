package models

import (
	"github.com/go-redis/redis"
	"os"
)

var Client *redis.Client

func NewClient() *redis.Client {
	var redisUrl string
	if redisUrl = os.Getenv("REDIS_URL"); redisUrl == "" {
		redisUrl = "redis://127.0.0.1:6379"
	}
	options, _ := redis.ParseURL(redisUrl)
	client := redis.NewClient(options)
	return client
}

func init() {
	Client = NewClient()
}
