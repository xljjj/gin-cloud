package model

import (
	"CloudDrive/mysql"
	"CloudDrive/util"
	"path"
	"strconv"
	"strings"
	"time"
)

// MyFile 文件
type MyFile struct {
	Id             int
	FileName       string //文件名
	FileHash       string //文件哈希值
	FileStoreId    int    //文件仓库id
	FilePath       string //文件存储路径
	DownloadNum    int    //下载次数
	UploadTime     string //上传时间
	ParentFolderId int    //父文件夹id
	Size           int64  //文件大小
	SizeStr        string //文件大小单位
	Type           int    //文件类型
	Postfix        string //文件后缀
}

// CreateFile 创建文件
func CreateFile(fileFullName string, fileHash string, fileSize int64, fId int, fileStoreId int) {
	//后缀
	fileSuffix := path.Ext(fileFullName)
	//文件名
	fileName := fileFullName[0 : len(fileFullName)-len(fileSuffix)]

	var sizeStr string
	if fileSize < 1048576 {
		sizeStr = strconv.FormatInt(fileSize/1024, 10) + "KB"
	} else {
		sizeStr = strconv.FormatInt(fileSize/102400, 10) + "MB"
	}

	myFile := MyFile{
		FileName:       fileName,
		FileHash:       fileHash,
		FileStoreId:    fileStoreId,
		FilePath:       "",
		DownloadNum:    0,
		UploadTime:     time.Now().Format("2006-01-02 15:04:05"),
		ParentFolderId: fId,
		Size:           fileSize / 1024,
		SizeStr:        sizeStr,
		Type:           util.GetFileTypeInt(fileSuffix),
		Postfix:        strings.ToLower(fileSuffix),
	}

	mysql.DB.Create(&myFile)
}

// GetUserFiles 获取用户所有文件
func GetUserFiles(parentId int, storeId int) (files []MyFile) {
	mysql.DB.Find(&files, "file_store_id = ? and parent_folder_id = ?", storeId, parentId)
	return files
}

// ReduceStoreSize 上传文件后减去可用容量
func ReduceStoreSize(size int64, storeId int) {
	var fileStore FileStore
	mysql.DB.First(&fileStore, storeId)
	fileStore.CurrentSize = fileStore.CurrentSize + size/1024
	fileStore.MaxSize = fileStore.MaxSize - size/1024
	mysql.DB.Save(&fileStore)
}

// GetUserFileNum 得到用户文件数量
func GetUserFileNum(storeId int) (num int64) {
	var files []MyFile
	mysql.DB.Find(&files, "file_store_id = ?", storeId).Count(&num)
	return num
}

// GetFileDetailUse 获取用户文件使用明细情况
func GetFileDetailUse(fileStoreId int) map[string]int64 {
	var files []MyFile
	var (
		docCount   int64
		imgCount   int64
		videoCount int64
		musicCount int64
		otherCount int64
	)

	fileDetailUseMap := make(map[string]int64, 0)

	//文档类型
	docCount = mysql.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, 1).RowsAffected
	fileDetailUseMap["docCount"] = docCount
	////图片类型
	imgCount = mysql.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, 2).RowsAffected
	fileDetailUseMap["imgCount"] = imgCount
	//视频类型
	videoCount = mysql.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, 3).RowsAffected
	fileDetailUseMap["videoCount"] = videoCount
	//音乐类型
	musicCount = mysql.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, 4).RowsAffected
	fileDetailUseMap["musicCount"] = musicCount
	//其他类型
	otherCount = mysql.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, 5).RowsAffected
	fileDetailUseMap["otherCount"] = otherCount

	return fileDetailUseMap
}

// GetFilesByType 根据文件类型获取文件
func GetFilesByType(fileType int, fileStoreId int) (files []MyFile) {
	mysql.DB.Find(&files, "file_store_id = ? and type = ?", fileStoreId, fileType)
	return files
}

// FileExist 判断文件是否已存在
func FileExist(fId int, fileFullName string) bool {
	var file MyFile
	//获取文件后缀
	fileSuffix := strings.ToLower(path.Ext(fileFullName))
	//获取文件名
	filePrefix := fileFullName[0 : len(fileFullName)-len(fileSuffix)]

	mysql.DB.Find(&file, "parent_folder_id = ? and file_name = ? and postfix = ?", fId, filePrefix, fileSuffix)

	return file.Size != 0
}

// FileOssExist 通过hash判断文件是否已上传过oss
func FileOssExist(fileHash string) bool {
	var file MyFile
	mysql.DB.Find(&file, "file_hash = ?", fileHash)
	return file.FileHash != ""
}

// GetFileById 通过fileId获取文件信息
func GetFileById(fId string) (file MyFile) {
	mysql.DB.First(&file, fId)
	return file
}

// DownloadNumAdd 文件下载次数+1
func DownloadNumAdd(fId int) {
	var file MyFile
	mysql.DB.First(&file, fId)
	file.DownloadNum = file.DownloadNum + 1
	mysql.DB.Save(&file)
}

// DeleteUserFile 删除数据库中的文件
func DeleteUserFile(fId int, folderId int, storeId int) {
	mysql.DB.Where("id = ? and file_store_id = ? and parent_folder_id = ?", fId, storeId, folderId).Delete(MyFile{})
}
