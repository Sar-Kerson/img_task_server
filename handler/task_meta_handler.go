package handler

import (
	"log"

	"github.com/Sar-Kerson/img_task_server/model"
	"github.com/gin-gonic/gin"
)

func GetTaskMetaHandler(c *gin.Context) {
	tid := c.DefaultQuery("task_id", "")
	if tid == "" {
		RespData(c, nil)
	}
	meta, err := model.GetTaskMeta(tid)
	if err != nil {
		log.Printf("[GetTaskMetaHandler] GetTaskMeta failed, err: %s", err.Error())
		RespErr(c, err)
	}
	RespData(c, meta)
}
