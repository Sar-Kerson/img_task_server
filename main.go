package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sar-Kerson/img_task_server/logger"
	"github.com/Sar-Kerson/img_task_server/handler"
	"github.com/Sar-Kerson/img_task_server/dal"
	"github.com/gin-gonic/gin"
)

var (
	srv *http.Server
)

func init() {
	router := getRouter()
	srv = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
}

// 释放用户资源
func shutdown() {
	dal.Close()
	logger.Close()
	return
}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/upload", handler.UploadHandler)
	router.GET("/task", handler.GetTaskMetaHandler)
	router.GET("/my/tasks", handler.GetUserTaskList)
	return router
}

func main() {
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// release customized resources
	defer shutdown()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
