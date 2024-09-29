package controller

import (
	"CloudDrive/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

// AddFolder 新建文件夹
func AddFolder(c *gin.Context) {
	userNameAny, _ := c.Get("userName")
	userName := fmt.Sprintf("%v", userNameAny)
	//获取用户信息
	user := model.FindSimpleUserByUserName(userName)

	// 用于接收 JSON 请求体的数据结构
	var requestData struct {
		FolderName string `json:"folderName"`
		ParentId   string `json:"parentId"`
	}

	// 解析 JSON 请求体
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求"})
		return
	}

	folderName := requestData.FolderName
	parentId, _ := strconv.Atoi(requestData.ParentId)

	if model.FolderNameExists(parentId, folderName) {
		c.JSON(http.StatusForbidden, gin.H{"error": "文件名已存在！"})
		return
	}

	//新建文件夹数据
	folder := model.CreateFileFolder(folderName, parentId, user.FileStoreId)

	//创建路径
	folderPath := "./file" + model.GetFolderPath(folder)
	_ = os.MkdirAll(folderPath, os.ModePerm)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

// UpdateFolder 修改文件夹名
func UpdateFolder(c *gin.Context) {
	// 用于接收 JSON 请求体的数据结构
	var requestData struct {
		FolderName string `json:"folderName"`
		FolderId   string `json:"folderId"`
		ParentId   string `json:"parentId"`
	}

	// 解析 JSON 请求体
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求"})
		return
	}

	folderName := requestData.FolderName
	folderId, _ := strconv.Atoi(requestData.FolderId)
	parentId, _ := strconv.Atoi(requestData.ParentId)

	folder := model.GetFolderById(folderId)

	if folder.FileFolderName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件夹不存在"})
		return
	}

	if model.FolderNameExists(parentId, folderName) {
		c.JSON(http.StatusForbidden, gin.H{"error": "新文件名已存在！"})
		return
	}

	model.UpdateFolderName(folderId, folderName)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

// DeleteFolder 删除文件夹
func DeleteFolder(c *gin.Context) {
	folderIdStr := c.Query("fId")
	if folderIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件夹请求不存在"})
		return
	}
	folderId, _ := strconv.Atoi(folderIdStr)
	folder := model.GetFolderById(folderId)
	if folder.FileFolderName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件夹不存在"})
		return
	}
	//删除文件夹路径下所有内容
	folderPath := "./file" + model.GetFolderPath(folder)
	err := os.RemoveAll(folderPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "删除文件夹失败",
		})
		return
	}
	model.DeleteFileFolder(folderId)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
