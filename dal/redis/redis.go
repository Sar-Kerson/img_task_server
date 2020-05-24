package redis

import (
	"github.com/go-redis/redis/v8"
)


var (
	Client *redis.ClusterClient
)

func init() {
	Client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:6380",
			"127.0.0.1:6381",
			"127.0.0.1:6382",
			"127.0.0.1:6383",
			"127.0.0.1:6384",
			"127.0.0.1:6385",
			"127.0.0.1:6386",
		},
	})
}



func Close() {
	_ = Client.Close()
}