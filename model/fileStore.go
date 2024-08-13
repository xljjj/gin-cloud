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

// GetFileStoreByUserId 根据用户ID得到文件仓库
func GetFileStoreByUserId(userId int) (fileStore FileStore) {
	mysql.DB.Find(&fileStore, "user_id = ?", userId)
	return fileStore
}

// CheckCapacity 检查用户仓库的容量
func CheckCapacity(storeId int, fileSize int64) bool {
	var fileStore FileStore
	mysql.DB.First(&fileStore, "id = ?", storeId)
	return fileStore.MaxSize >= fileSize
}
