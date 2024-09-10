package controller

import (
	"CloudDrive/model"
	"CloudDrive/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	fIdStr := c.Query("fId")
	fId, _ := strconv.Atoi(fIdStr)

	file := model.GetFileById(fId)
	if file.FileHash == "" {
		return
	}

	//从oss获取文件
	fileData := util.DownloadOss(file.FileHash, file.Suffix)
	//下载次数+1
	model.DownloadNumAdd(fId)

	c.Header("Content-disposition", "attachment;filename=\""+file.FileName+file.Suffix+"\"")
	c.Data(http.StatusOK, "application/octect-stream", fileData)
}

// DeleteFile 删除文件
func DeleteFile(c *gin.Context) {
	fId := c.DefaultQuery("fId", "")
	folderId := c.Query("folder")
	if fId == "" {
		return
	}

	//删除数据库文件数据
	fIdInt, _ := strconv.Atoi(fId)
	model.DeleteFileById(fIdInt)

	c.Redirect(http.StatusMovedPermanently, "/cloud/file?fid="+folderId)
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
