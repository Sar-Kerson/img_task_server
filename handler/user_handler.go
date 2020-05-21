package handler

import (
	"log"

	"github.com/Sar-Kerson/img_task_server/model"
	"github.com/gin-gonic/gin"
)

func GetUserTaskList(c *gin.Context) {
	uid := c.DefaultQuery("user_id", "")
	if uid == "" {
		RespData(c, nil)
		return
	}

	tids, err := model.GetUserTaskIDList(uid)
	if err != nil {
		RespErr(c, err)
		log.Printf("[GetUserTaskList] GetUserTaskIDList, err: %v", err)
		return
	}

	taskList, err := model.MGetUserTaskList(tids)
	if err != nil {
		RespErr(c, err)
		log.Printf("[GetUserTaskList] MGetUserTaskList, err: %v", err)
		return
	}

	RespData(c, taskList)
}
