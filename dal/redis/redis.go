package redis

import "github.com/go-redis/redis/v7"

var (
	Client *redis.Client
)

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6382",
		Password: "",
		DB:       0,
	})
}

func Close() {
	_ = Client.Close()
}