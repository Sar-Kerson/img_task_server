package dal

import (
	"github.com/Sar-Kerson/img_task_server/dal/kafka"
	"github.com/Sar-Kerson/img_task_server/dal/redis"
)

func init() {
}

func Close() {
	kafka.Close()
	redis.Close()
}
