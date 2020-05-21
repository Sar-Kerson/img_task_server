package redis

import "testing"

func TestRedis(t *testing.T) {
	pong, err := Client.Ping().Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf(pong)
}
