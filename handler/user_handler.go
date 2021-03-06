package handler

import (
	"log"

	"github.com/Sar-Kerson/img_task_server/model"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	uid := c.DefaultQuery("uid", "")
	passwd := c.DefaultQuery("passwd", "")
	if uid == "" || passwd == "" {
		RespData(c, false)
		return
	}
	err := model.ValidatePassword(c,uid, passwd)
	if err != nil {
		RespData(c, false)
		log.Printf("[UploadHandler] ValidatePassword, err: %v", err)
		return
	}
	RespData(c, true)
}

func SignupHandler(c *gin.Context) {
	uid := c.DefaultQuery("uid", "")
	passwd := c.DefaultQuery("passwd", "")
	if uid == "" || passwd == "" {
		RespData(c, false)
		return
	}
	exist, _ := model.CheckUserExist(c,uid)
	if exist {
		RespData(c, false)
		log.Printf("[SignupHandler] user exist")
		return
	}
	err := model.SetUserInfo(c,uid, passwd)
	if err != nil {
		RespErr(c, err)
		log.Printf("[SignupHandler] SetUserInfo, err: %v", err)
		return
	}
	RespData(c, true)
}

func GetUserTaskList(c *gin.Context) {
	uid := c.DefaultQuery("user_id", "")
	if uid == "" {
		RespData(c, nil)
		return
	}

	tids, err := model.GetUserTaskIDList(c, uid)
	if err != nil {
		RespErr(c, err)
		log.Printf("[GetUserTaskList] GetUserTaskIDList, err: %v", err)
		return
	}

	taskList, err := model.MGetUserTaskList(c,tids)
	res := make([]model.TaskApiData, 0, len(taskList))
	for _, meta := range taskList {
		res = append(res, *model.NewTaskApiDataFromTaskMeta(meta))
	}
	if err != nil {
		RespErr(c, err)
		log.Printf("[GetUserTaskList] MGetUserTaskList, err: %v", err)
		return
	}

	RespData(c, res)
}
