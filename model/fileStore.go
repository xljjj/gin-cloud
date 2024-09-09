package model

import (
	"CloudDrive/mysql"
)

// FileStore 文件仓库
type FileStore struct {
	Id          int
	UserId      int
	CurrentSize int64
	MaxSize     int64
}

func (FileStore) TableName() string {
	return "file_store"
}

// GetFileStoreByUserId 根据用户ID得到文件仓库
func GetFileStoreByUserId(userId int) (fileStore FileStore) {
	mysql.DB.Find(&fileStore, "user_id = ?", userId)
	return fileStore
}

// CheckCapacity 检查用户仓库的容量
func CheckCapacity(storeId int, fileSize int64) bool {
	var fileStore FileStore
	mysql.DB.First(&fileStore, "id = ?", storeId)
	return fileStore.MaxSize-fileStore.CurrentSize >= fileSize
}

// AddStoreSize 上传文件后增加现有容量
func AddStoreSize(size int64, storeId int) {
	var fileStore FileStore
	mysql.DB.First(&fileStore, storeId)
	fileStore.CurrentSize = fileStore.CurrentSize + size/1024
	mysql.DB.Save(&fileStore)
}
