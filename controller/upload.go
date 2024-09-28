package controller

import (
	"CloudDrive/model"
	"CloudDrive/redis"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Upload 上传文件页面
func Upload(c *gin.Context) {
	userNameAny, _ := c.Get("userName")
	userName := fmt.Sprintf("%v", userNameAny)
	//获取用户信息
	user := model.FindSimpleUserByUserName(userName)
	// 获取当前文件夹ID，根目录为0
	fIdStr := c.DefaultQuery("fId", "0")
	fId, _ := strconv.Atoi(fIdStr)
	fmt.Println("folder:", fId)
	// 获取当前目录信息
	folder := model.GetFolderById(fId)

	c.HTML(http.StatusOK, "upload.html", gin.H{
		"user":   user,
		"fId":    fId,
		"folder": folder,
	})
}

// HandleUpload 处理上传文件
func HandleUpload(c *gin.Context) {
	userNameAny, _ := c.Get("userName")
	userName := fmt.Sprintf("%v", userNameAny)
	//获取用户信息
	user := model.FindSimpleUserByUserName(userName)

	fIdStr := c.GetHeader("fid")
	fId, _ := strconv.Atoi(fIdStr)

	// 接收上传文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "未接收到文件"})
		return
	}
	defer file.Close() // 确保文件在使用后被关闭

	// 获取文件名和文件长度
	fileName := header.Filename
	fileSize := header.Size
	KBSize := (fileSize + 1023) / 1024 //为方便处理，统一单位为KB，不足1KB当成1KB

	//判断当前文件夹是否有同名文件
	if model.FileExist(fId, fileName) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "同名文件已存在！",
		})
		return
	}

	//判断用户的容量是否足够
	if ok := model.CheckCapacity(user.FileStoreId, KBSize); !ok {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "用户剩余容量不足！",
		})
		return
	}

	// 新建文件
	var filePath string
	if fId == 0 {
		filePath = "./file/" + strconv.Itoa(user.FileStoreId) + "/" + fileName
	} else {
		folder := model.GetFolderById(fId)
		filePath = "./file" + model.GetFolderPath(folder) + "/" + fileName
	}

	newFile, err := os.Create(filePath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法新建文件",
		})
		return
	}
	defer newFile.Close()
	_, err = io.Copy(newFile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法新建文件",
		})
		return
	}
	myFile := model.CreateFile(fileName, "", KBSize, fId, user.FileStoreId)

	// 上传成功后增加现有容量
	model.AddStoreSize(KBSize, user.FileStoreId)

	// 如果文件较小，将文件存入Redis
	if KBSize <= 100 {
		//重新获得文件流
		file, _, _ = c.Request.FormFile("file")
		fileContent, _ := io.ReadAll(file)
		// 新建一个带有超时的上下文
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel() // 操作完成后调用取消函数，避免资源泄漏
		_ = redis.SetKey(ctx, strconv.Itoa(myFile.Id), fileContent, time.Hour)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
