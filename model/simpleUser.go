package model

import (
	"CloudDrive/mysql"
	"gorm.io/gorm"
	"time"
)

// SimpleUser 自建用户
type SimpleUser struct {
	Id            int
	UserName      string
	Password      string
	NickName      string
	Ext           string
	LastLoginTime time.Time
	FileStoreId   int
}

func (SimpleUser) TableName() string {
	return "simple_user"
}

// CreateSimpleUser 创建自建用户
func CreateSimpleUser(s *SimpleUser) *gorm.DB {
	return mysql.DB.Create(s)
}

// DeleteSimpleUser 删除自建用户
func DeleteSimpleUser(s *SimpleUser) *gorm.DB {
	return mysql.DB.Delete(s)
}

// UpdateSimpleUser 更新自建用户
func UpdateSimpleUser(s *SimpleUser) *gorm.DB {
	return mysql.DB.Save(s)
}

// FindSimpleUserByUserName 查找用户
func FindSimpleUserByUserName(name string) SimpleUser {
	var simpleUser SimpleUser
	_ = mysql.DB.Where("user_name = ?", name).First(&simpleUser)
	return simpleUser
}
