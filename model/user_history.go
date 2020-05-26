package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	redis_util "github.com/Sar-Kerson/img_task_server/dal/redis"
	"github.com/go-redis/redis/v8"
)

func InsertToUserTaskList(ctx context.Context, uid, taskId string) error {
	log.Printf("[InsertToUserTaskList] get req, uid: %s, taskId: %s", uid, taskId)
	key := fmtUserTaskListKey(uid)
	if err := redis_util.Client.ZAdd(ctx, key, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: taskId,
	}).Err(); err != nil {
		log.Printf("[InsertToUserTaskList] zadd failed, key: %s, err: %s", key, err.Error())
		return err
	}
	return nil
}

func GetUserTaskIDList(ctx context.Context, uid string) ([]string, error) {
	log.Printf("[GetUserTaskIDList] get req, uid: %s", uid)
	key := fmtUserTaskListKey(uid)
	vals, err := redis_util.Client.ZRevRange(ctx, key, 0, -1).Result()
	if err != nil {
		log.Printf("[GetUserTaskIDList] ZRevRange failed, key: %s, err: %s", key, err.Error())
		return []string{}, err
	}
	return vals, err
}

func MGetUserTaskList(ctx context.Context, tids []string) ([]TaskMeta, error) {
	if len(tids) ==  0 {
		return []TaskMeta{}, nil
	}

	res := make([]TaskMeta, 0, len(tids))

	keys := make([]string, 0, len(tids))
	for _, tid := range tids {
		keys = append(keys, fmtTaskMetaKey(tid))
	}

	valStrs, err := redis_util.Client.MGet(ctx, keys...).Result()
	if err != nil {
		log.Printf("[MGetUserTaskList] MGet failed, keys: %v, err: %s", keys, err.Error())
		return []TaskMeta{}, err
	}

	for _, valStr := range valStrs {
		if valStr == nil {
			continue
		}
		meta := TaskMeta{}
		err := json.Unmarshal([]byte(valStr.(string)), &meta)
		if err != nil {
			log.Printf("[MGetUserTaskList] Unmarshal failed, valStr: %s, err: %s", valStr, err.Error())
			continue
		}
		res = append(res, meta)
	}

	return res, nil
}

func fmtUserTaskListKey(uid string) string {
	return fmt.Sprintf("u:l:%s", uid)
}
