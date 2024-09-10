package controller

import (
	"CloudDrive/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	userNameAny, _ := c.Get("userName")
	userName := fmt.Sprintf("%v", userNameAny)
	//获取用户信息
	user := model.FindSimpleUserByUserName(userName)
	//获取用户仓库信息
	userFileStore := model.GetFileStoreByUserId(user.Id)
	//获取用户文件数量
	fileCount := model.GetUserFileNum(user.FileStoreId)
	//获取用户文件夹数量
	fileFolderCount := model.GetUserFolderNum(user.FileStoreId)
	//获取用户文件使用明细数量
	fileDetailUse := model.GetFileDetailUse(user.FileStoreId)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"user":            user,
		"userFileStore":   userFileStore,
		"fileCount":       fileCount,
		"fileFolderCount": fileFolderCount,
		"fileDetailUse":   fileDetailUse,
	})
}
