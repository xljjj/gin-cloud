package controller

import (
	"CloudDrive/model"
	"CloudDrive/redis"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

// File 文件信息
func File(c *gin.Context) {
	userNameAny, _ := c.Get("userName")
	userName := fmt.Sprintf("%v", userNameAny)
	//获取用户信息
	user := model.FindSimpleUserByUserName(userName)
	// 获取当前文件夹ID，根目录为0
	fIdStr := c.DefaultQuery("fId", "0")
	fId, _ := strconv.Atoi(fIdStr)
	// 获取当前目录信息
	folder := model.GetFolderById(fId)

	//获取当前目录所有文件
	files := model.GetFolderFiles(fId, user.FileStoreId)

	c.HTML(http.StatusOK, "file.html", gin.H{
		"user":   user,
		"fId":    fId,
		"folder": folder,
		"files":  files,
	})
}

// AddFolder 新建文件夹
func AddFolder(c *gin.Context) {
	openId, _ := c.Get("openId")
	user := model.GetUser(fmt.Sprintf("%v", openId))

	folderName := c.PostForm("fileFolderName")
	parentIdStr := c.DefaultPostForm("parentFolderId", "0")
	parentId, _ := strconv.Atoi(parentIdStr)

	//新建文件夹数据
	model.CreateFileFolder(folderName, parentId, user.FileStoreId)

	//获取父文件夹信息
	parent := model.GetFolderById(parentId)

	c.Redirect(http.StatusMovedPermanently, "/cloud/file?fId="+parentIdStr+"&fName="+parent.FileFolderName)
}

// DownloadFile 下载文件
func DownloadFile(c *gin.Context) {
	//fIdStr := c.DefaultQuery("fId","0")  //文件夹ID
	//fId, _ := strconv.Atoi(fIdStr)
	fileIdStr := c.Query("fileId")
	if fileIdStr == "" {
		return
	}
	fileId, _ := strconv.Atoi(fileIdStr)
	file := model.GetFileById(fileId)
	if file.FileName == "" {
		return
	}
	// Redis中有文件数据则从Redis中取出
	// 新建一个带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel() // 操作完成后调用取消函数，避免资源泄漏
	flag, err := redis.KeyExists(ctx, fileIdStr)
	if flag && err == nil {
		fileData, _ := redis.GetKey(ctx, fileIdStr)
		// 设置 Content-Disposition 头，提示浏览器下载文件
		c.Header("Content-Disposition", "attachment; filename="+file.FileName+file.Suffix)
		c.Header("Content-Type", "application/octet-stream")
		// 返回文件内容
		c.Data(http.StatusOK, "application/octet-stream", []byte(fileData))
	} else {
		// 根据 fileId 确定文件路径
		filePath := "./file/" + strconv.Itoa(file.FileStoreId) + "/" + file.FileName + file.Suffix
		// 设置 Content-Disposition 头，提示浏览器下载文件
		c.Header("Content-Disposition", "attachment; filename="+file.FileName+file.Suffix)
		c.Header("Content-Type", "application/octet-stream")
		// 下载次数+1
		model.DownloadNumAdd(fileId)
		// 返回文件内容
		c.File(filePath)
	}
}

// DeleteFile 删除文件
func DeleteFile(c *gin.Context) {
	//fIdStr := c.DefaultQuery("fId","0")  //文件夹ID
	//fId, _ := strconv.Atoi(fIdStr)
	fileIdStr := c.GetHeader("fileId")
	if fileIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件请求不存在"})
		return
	}
	fileId, _ := strconv.Atoi(fileIdStr)
	file := model.GetFileById(fileId)
	if file.FileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件存储不存在"})
		return
	}
	// 删除文件存储
	err := os.Remove("./file/" + strconv.Itoa(file.FileStoreId) + "/" + file.FileName + file.Suffix)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "删除文件失败",
		})
		return
	}
	// 加回用户可用容量
	model.AddStoreSize(-file.Size, file.FileStoreId)
	// 删除数据库文件数据
	model.DeleteFileById(fileId)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

// DeleteFileFolder 删除文件夹
func DeleteFileFolder(c *gin.Context) {
	fId := c.DefaultQuery("fId", "")
	if fId == "" {
		return
	}

	//获取要删除的文件夹信息 取到父级目录重定向
	fIdInt, _ := strconv.Atoi(fId)
	folderInfo := model.GetFolderById(fIdInt)

	//删除文件夹并删除文件夹中的文件信息
	model.DeleteFileFolder(fIdInt)

	c.Redirect(http.StatusMovedPermanently, "/cloud/file?fId="+strconv.Itoa(folderInfo.ParentFolderId))
}

// UpdateFileFolder 修改文件夹名
func UpdateFileFolder(c *gin.Context) {
	fileFolderName := c.PostForm("fileFolderName")
	fileFolderId := c.PostForm("fileFolderId")

	fileFolderIdInt, _ := strconv.Atoi(fileFolderId)
	fileFolder := model.GetFolderById(fileFolderIdInt)

	model.UpdateFolderName(fileFolderIdInt, fileFolderName)

	c.Redirect(http.StatusMovedPermanently, "/cloud/file?fId="+strconv.Itoa(fileFolder.ParentFolderId))
}
