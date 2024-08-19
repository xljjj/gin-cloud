package controller

import (
	"CloudDrive/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	openId, _ := c.Get("openId")
	//获取用户信息
	user := model.GetUser(fmt.Sprintf("%v", openId))
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
		"curIndex":        "active",
		"userFileStore":   userFileStore,
		"fileCount":       fileCount,
		"fileFolderCount": fileFolderCount,
		"fileDetailUse":   fileDetailUse,
	})
}
