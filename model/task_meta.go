package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Sar-Kerson/img_task_server/dal/redis"
)

const (
	TASK_STATUS_PROCESSING = 1
	TASK_STATUS_SUC        = 2
	TASK_STATUS_FAILED     = 10
)

var (
	ERR_INVALID_PARAMS = errors.New("invalid params")
)

type TaskMeta struct {
	TaskID     string `json:"task_id"`
	UserID     string `json:"user_id"`
	CreateTime int64  `json:"create_time"`
	ProcStatus int64  `json:"proc_status"`
	InputURL   string `json:"input_url"`
	OutputURL  string `json:"output_url"`
}

func NewTaskMeta(tid, uid, inUrl string) *TaskMeta {
	return &TaskMeta{
		TaskID:     tid,
		UserID:     uid,
		CreateTime: time.Now().Unix(),
		ProcStatus: TASK_STATUS_PROCESSING,
		InputURL:   inUrl,
	}
}

func SetTaskMeta(taskMeta *TaskMeta) error {
	if taskMeta == nil {
		return ERR_INVALID_PARAMS
	}
	key := fmtTaskMetaKey(taskMeta.TaskID)
	val, err := json.Marshal(taskMeta)
	if err != nil {
		log.Printf("[SetTaskMeta] Marshal failed, key: %s, err: %s", key, err.Error())
		return err
	}
	return redis.Client.Set(key, val, 0).Err()
}

func GetTaskMeta(taskId string) (*TaskMeta, error) {
	key := fmtTaskMetaKey(taskId)
	valStr, err := redis.Client.Get(key).Result()
	if err != nil {
		log.Printf("[GetTaskMeta] Get(key) failed, key: %s, err: %s", key, err.Error())
		return nil, err
	}
	meta := &TaskMeta{}
	err = json.Unmarshal([]byte(valStr), meta)
	if err != nil {
		log.Printf("[GetTaskMeta] Unmarshal failed, err: %s", err.Error())
		return nil, err
	}
	return meta, nil
}

func fmtTaskMetaKey(taskId string) string {
	return fmt.Sprintf("test:%s", taskId)
}
