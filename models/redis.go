package redis

import (
	"github.com/go-redis/redis"
	"os"
	"strings"
)

var Client *redis.Client

func NewClient(host string, port string) *redis.Client {
	addr := []string{host, port}
	client := redis.NewClient(&redis.Options{
		Addr:     strings.Join(addr, ":"),
		Password: "",
		DB:       0,
	})
	return client
}

func init() {
	var redisHost, redisPort string
	redisHost = os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "127.0.0.1"
	}

	redisPort = os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}
	Client = NewClient(redisHost, redisPort)
}
