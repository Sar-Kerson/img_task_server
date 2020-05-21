package redis

import "github.com/go-redis/redis/v7"

var (
	Client *redis.Client
)

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "35.240.132.243:6379",
		Password: "",
		DB:       0,
	})
}

func Close() {
	_ = Client.Close()
}