package model

import (
	"CloudDrive/mysql"
	"CloudDrive/util"
	"path"
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
	Size           int64  //文件大小（默认单位KB）
	Type           int    //文件类型
	Suffix         string //文件后缀
}

func (MyFile) TableName() string {
	return "my_file"
}

// CreateFile 创建文件
func CreateFile(fileFullName string, fileHash string, fileSize int64, fId int, fileStoreId int) {
	//后缀
	fileSuffix := path.Ext(fileFullName)
	//文件名
	fileName := fileFullName[0 : len(fileFullName)-len(fileSuffix)]

	myFile := MyFile{
		FileName:       fileName,
		FileHash:       fileHash,
		FileStoreId:    fileStoreId,
		FilePath:       "",
		DownloadNum:    0,
		UploadTime:     time.Now().Format("2006-01-02 15:04:05"),
		ParentFolderId: fId,
		Size:           fileSize / 1024,
		Type:           util.GetFileTypeInt(fileSuffix),
		Suffix:         strings.ToLower(fileSuffix),
	}

	mysql.DB.Create(&myFile)
}

// GetFolderFiles 获取文件夹下所有文件
func GetFolderFiles(parentId int, storeId int) (files []MyFile) {
	mysql.DB.Find(&files, "file_store_id = ? and parent_folder_id = ?", storeId, parentId)
	return files
}

// GetUserFileNum 得到用户文件总数量
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
	fileName := fileFullName[0 : len(fileFullName)-len(fileSuffix)]

	mysql.DB.Find(&file, "parent_folder_id = ? and file_name = ? and suffix = ?", fId, fileName, fileSuffix)

	return file.FileName != ""
}

// FileOssExist 通过hash判断文件是否已上传过oss
func FileOssExist(fileHash string) bool {
	var file MyFile
	mysql.DB.Find(&file, "file_hash = ?", fileHash)
	return file.FileName != ""
}

// GetFileById 通过fileId获取文件信息
func GetFileById(fId int) (file MyFile) {
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
