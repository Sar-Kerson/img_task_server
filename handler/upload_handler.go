package handler

import (
	"bytes"
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"mime/multipart"

	"github.com/Sar-Kerson/img_task_server/dal/cloud_storage"
	"github.com/Sar-Kerson/img_task_server/model"
	"github.com/Sar-Kerson/img_task_server/util"
	"github.com/anthonynsimon/bild/transform"
	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	// TODO: token here
	uid := c.PostForm("user_id")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		RespErr(c, err)
		log.Printf("[UploadHandler] FormFile, err: %v", err)
		return
	}

	// 进行预处理操作，图像尺寸调整
	imgBytes, err := preprocessImage(file)
	if err != nil {
		RespErr(c, err)
		log.Printf("[UploadHandler] preprocessImage, err: %v", err)
		return
	}

	taskID := util.Md5Hash(util.MergeBytes(imgBytes, []byte(uid)))

	// 上传对象存储
	url, err := cloud_storage.Upload(c, imgBytes, taskID)
	if err != nil {
		RespErr(c, err)
		log.Printf("[UploadHandler] Upload, err: %v", err)
		return
	}

	// 先看下有没有元信息存在，有的话则不处理
	meta, err := model.GetTaskMeta(taskID)
	if err == nil && meta.ProcStatus != model.TASK_STATUS_FAILED {
		RespData(c, url)
		log.Printf("[UploadHandler] task already exist, taskID: %s", taskID)
		return
	}

	// 更新任务元信息
	err = updateTaskMeta(uid, taskID, url)
	if err != nil {
		RespErr(c, err)
		log.Printf("[UploadHandler] updateTaskMeta, err: %v", err)
		return
	}

	// 写入生产者
	err = model.CommitTask(taskID)
	if err != nil {
		RespErr(c, err)
		log.Printf("[UploadHandler] CommitTask, err: %v", err)
		return
	}

	RespData(c, url)
}

func preprocessImage(header *multipart.FileHeader) ([]byte, error) {
	// 打开文件
	fp, err := header.Open()
	if err != nil {
		log.Printf("[preprocessImage] header.Open, err: %v", err)
		return []byte{}, err
	}

	// 加载图像
	img, _, err := image.Decode(fp)
	if err != nil {
		log.Printf("[preprocessImage] image.Decode, err: %v", err)
		return []byte{}, err
	}

	// 图像预处理
	newImage := transform.Resize(img, 256, 256, transform.Linear)

	newBuf := new(bytes.Buffer)
	err = jpeg.Encode(newBuf, newImage, nil)
	if err != nil {
		log.Printf("[preprocessImage] jpeg.Encode, err: %v", err)
		return []byte{}, err
	}

	return newBuf.Bytes(), nil
}

func updateTaskMeta(uid, taskID, url string) error {
	// 生成元信息
	meta := model.NewTaskMeta(taskID, uid, url)

	// 插入用户历史提交
	err := model.InsertToUserTaskList(uid, taskID)
	if err != nil {
		log.Printf("[updateTaskMeta] InsertToUserTaskList failed, err: %v", err)
		return err
	}

	// 插入任务元信息
	err = model.SetTaskMeta(meta)
	if err != nil {
		log.Printf("[updateTaskMeta] SetTaskMeta failed, err: %v", err)
		return err
	}

	return nil
}
