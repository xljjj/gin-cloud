package model

import (
	"CloudDrive/mysql"
	"CloudDrive/util"
	"strings"
	"time"
)

// Share 分享
type Share struct {
	Id       int
	Code     string
	FileId   int
	Username string
	Hash     string
}

func (Share) TableName() string {
	return "share"
}

// CreateShare 创建分享
func CreateShare(code string, username string, fId int) string {
	share := Share{
		Code:     strings.ToLower(code),
		FileId:   fId,
		Username: username,
		Hash:     util.Md5Encode(code + string(time.Now().Unix())),
	}
	mysql.DB.Create(&share)
	return share.Hash
}

// GetShare 根据Hash查询分享
func GetShare(f string) (share Share) {
	mysql.DB.Find(&share, "hash = ?", f)
	return
}

// VerifyShare 校验分享
func VerifyShare(fId, code string) bool {
	var share Share
	mysql.DB.Find(&share, "file_id = ? and code = ?", fId, code)
	return share.Id != 0
}
