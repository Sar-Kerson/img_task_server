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
	meta, err := model.GetTaskMeta(c,tid)
	if err != nil {
		log.Printf("[GetTaskMetaHandler] GetTaskMeta failed, err: %s", err.Error())
		RespErr(c, err)
	}
	RespData(c, meta)
}

func TestCommitHandler(c *gin.Context)  {
	err := model.CommitTask("5b5f663b7a3a8f9e4f32dc62e8de848f")
	if err != nil {
		log.Printf("[TestCommitHandler] CommitTask failed, err: %s", err.Error())
		RespErr(c, err)
	}
	RespData(c, nil)
}