package redis

import (
	"github.com/go-redis/redis"
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
	Client = NewClient("127.0.0.1", "6379")
}
