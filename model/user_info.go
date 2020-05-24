package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	redis_util "github.com/Sar-Kerson/img_task_server/dal/redis"
	"github.com/go-redis/redis/v8"
)

type UserInfo struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

func ValidatePassword(ctx context.Context,uid, pwd string) error {
	key := fmtUserInfoKey(uid)
	valStr, err := redis_util.Client.Get(ctx, key).Result()
	if err != nil {
		log.Printf("[ValidatePassword] Get(key) failed, key: %s, err: %s", key, err.Error())
		return err
	}
	userInfo := &UserInfo{}
	err = json.Unmarshal([]byte(valStr), userInfo)
	if err != nil {
		log.Printf("[ValidatePassword] Unmarshal failed, err: %s", err.Error())
		return err
	}
	if userInfo.Password != pwd {
		return errors.New("invalid Password")
	}
	return nil
}

func CheckUserExist(ctx context.Context, uid string) (bool, error) {
	key := fmtUserInfoKey(uid)
	err := redis_util.Client.Get(ctx,key).Err()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func SetUserInfo(ctx context.Context, uid, pwd string) error {
	userInfo := UserInfo{uid, pwd}
	key := fmtUserInfoKey(uid)
	val, err := json.Marshal(userInfo)
	if err != nil {
		log.Printf("[SetTaskMeta] Marshal failed, key: %s, err: %s", key, err.Error())
		return err
	}
	return redis_util.Client.Set(ctx, key, val, 0).Err()
}

func fmtUserInfoKey(uid string) string {
	return fmt.Sprintf("u:%s", uid)
}
