package model

import (
	"CloudDrive/util"
	"time"
)

// User 用户
type User struct {
	Id           int
	OpenId       string
	FileStoreId  int
	UserName     string
	RegisterTime time.Time
	ImagePath    string
}

// CreateUser 创建用户和文件仓库
func CreateUser(openId string, userName string, imagePath string) {
	user := User{
		OpenId:       openId,
		FileStoreId:  0,
		UserName:     userName,
		RegisterTime: time.Now(),
		ImagePath:    imagePath,
	}
	util.DB.Create(&user)
	fileStore := FileStore{
		UserId:      user.Id,
		CurrentSize: 0,
		MaxSize:     1048576,
	}
	util.DB.Create(&fileStore)
	user.FileStoreId = fileStore.Id
	util.DB.Save(&user)
}

// UserExists 判断用户是否存在
func UserExists(openId string) bool {
	var user User
	util.DB.Find(&user, "open_id = ?", openId)
	return user.Id != 0
}

// GetUser 根据openID得到用户
func GetUser(openId string) (user User) {
	util.DB.Find(&user, "open_id = ?", openId)
	return user
}
