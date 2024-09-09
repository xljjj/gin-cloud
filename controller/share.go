package controller

import (
	"CloudDrive/model"
	"CloudDrive/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lifei6671/gocaptcha"
	"net/http"
	"strconv"
	"strings"
)

// ShareFile 创建分享文件
func ShareFile(c *gin.Context) {
	openId, _ := c.Get("openId")
	//获取用户信息
	user := model.GetUser(fmt.Sprintf("%v", openId))

	fId := c.Query("id")
	url := c.Query("url")
	//获取内容
	code := gocaptcha.RandText(4)

	fileId, _ := strconv.Atoi(fId)
	hash := model.CreateShare(code, user.UserName, fileId)

	c.JSON(http.StatusOK, gin.H{
		"url":  url + "?f=" + hash,
		"code": code,
	})
}

// SharePass 分享文件页面
func SharePass(c *gin.Context) {
	f := c.Query("f")

	//获取分享信息
	shareInfo := model.GetShare(f)
	//获取文件信息
	file := model.GetFileById(shareInfo.FileId)

	c.HTML(http.StatusOK, "share.html", gin.H{
		"id":       shareInfo.FileId,
		"username": shareInfo.UserName,
		"fileType": file.Type,
		"filename": file.FileName + file.Suffix,
		"hash":     shareInfo.Hash,
	})
}

// DownloadShareFile 下载分享文件
func DownloadShareFile(c *gin.Context) {
	fileIdStr := c.Query("id")
	code := c.Query("code")
	hash := c.Query("hash")
	fileId, _ := strconv.Atoi(fileIdStr)

	fileInfo := model.GetFileById(fileId)

	//校验提取码
	if ok := model.VerifyShare(fileId, strings.ToLower(code)); !ok {
		c.Redirect(http.StatusMovedPermanently, "/file/share?f="+hash)
		return
	}

	//从oss获取文件
	fileData := util.DownloadOss(fileInfo.FileHash, fileInfo.Suffix)
	//下载次数+1
	model.DownloadNumAdd(fileId)

	c.Header("Content-disposition", "attachment;filename=\""+fileInfo.FileName+fileInfo.Suffix+"\"")
	c.Data(http.StatusOK, "application/octect-stream", fileData)
}
