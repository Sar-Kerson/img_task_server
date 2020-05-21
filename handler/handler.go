package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RespData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":    0,
		"message":   "success",
		"data":      data,
		"timestamp": time.Now().Unix(),
	})
	log.Printf("[RespData] path: %v, data: %+v", c.Request.URL.Path, data)
}

func RespErr(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":    0,
		"message":   err.Error(),
		"data":      nil,
		"timestamp": time.Now().Unix(),
	})
	log.Printf("[RespErr] path: %s, err: %s", c.Request.URL.Path, err.Error())
}
