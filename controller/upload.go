package controller

import (
	"CloudDrive/model"
	"CloudDrive/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"strconv"
)

// Upload 上传文件页面
func Upload(c *gin.Context) {
	openId, _ := c.Get("openId")
	fId := c.DefaultQuery("fId", "0")
	//获取用户信息
	user := model.GetUser(fmt.Sprintf("%v", openId))
	//获取当前目录信息
	fIdInt, _ := strconv.Atoi(fId)
	currentFolder := model.GetFolderById(fIdInt)
	//获取当前目录所有的文件夹信息
	fileFolders := model.GetChildrenFolders(fIdInt, user.FileStoreId)
	//获取父级的文件夹信息
	parentFolder := model.GetFolderById(fIdInt)
	//获取当前目录所有父级
	currentAllParent := model.GetPathParents(parentFolder)
	//获取用户文件使用明细数量
	fileDetailUse := model.GetFileDetailUse(user.FileStoreId)

	c.HTML(http.StatusOK, "upload.html", gin.H{
		"user":             user,
		"currUpload":       "active",
		"fId":              currentFolder.Id,
		"fName":            currentFolder.FileFolderName,
		"fileFolders":      fileFolders,
		"parentFolder":     parentFolder,
		"currentAllParent": currentAllParent,
		"fileDetailUse":    fileDetailUse,
	})
}

// HandleUpload 处理上传文件
func HandleUpload(c *gin.Context) {
	openId, _ := c.Get("openId")
	//获取用户信息
	user := model.GetUser(fmt.Sprintf("%v", openId))

	FId := c.GetHeader("id")
	FIdInt, _ := strconv.Atoi(FId)

	//接收上传文件
	file, head, err := c.Request.FormFile("file")

	//判断当前文件夹是否有同名文件
	if ok := model.FileExist(FIdInt, head.Filename); !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 501,
		})
		return
	}

	//判断用户的容量是否足够
	if ok := model.CheckCapacity(int(head.Size), int64(user.FileStoreId)); ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 503,
		})
		return
	}

	if err != nil {
		fmt.Println("File Upload Error:", err.Error())
		return
	}
	defer file.Close()

	//文件保存本地的路径
	location := viper.GetString("update.location") + head.Filename

	//在本地创建一个新的文件
	newFile, err := os.Create(location)
	if err != nil {
		fmt.Println("File Create Error:", err.Error())
		return
	}
	defer newFile.Close()

	//将上传文件拷贝至新创建的文件中
	fileSize, err := io.Copy(newFile, file)
	if err != nil {
		fmt.Println("File Copy Error:", err.Error())
		return
	}

	//将光标移至开头
	_, _ = newFile.Seek(0, 0)
	fileHash := util.SHA256HashCode(newFile)

	//通过hash判断文件是否已上传过oss
	if ok := model.FileOssExist(fileHash); !ok {
		//上传至阿里云oss
		go util.UploadOss(head.Filename, fileHash)
	}
	//新建文件信息
	model.CreateFile(head.Filename, fileHash, fileSize, FIdInt, user.FileStoreId)
	//上传成功减去相应剩余容量
	model.ReduceStoreSize(fileSize/1024, user.FileStoreId)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
