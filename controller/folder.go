package controller

import (
	"CloudDrive/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
