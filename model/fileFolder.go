package model

import (
	"CloudDrive/mysql"
	"strconv"
	"time"
)

// FileFolder 文件夹
type FileFolder struct {
	Id             int
	FileFolderName string
	ParentFolderId int
	FileStoreId    int
	Time           string
}

func (FileFolder) TableName() string {
	return "file_folder"
}

// CreateFileFolder 新建文件夹
func CreateFileFolder(folderName string, parentId int, fileStoreId int) FileFolder {
	fileFolder := FileFolder{
		FileFolderName: folderName,
		ParentFolderId: parentId,
		FileStoreId:    fileStoreId,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
	}
	mysql.DB.Create(&fileFolder)
	return fileFolder
}

// GetFolderById 根据ID得到文件夹
func GetFolderById(fid int) (fileFolder FileFolder) {
	mysql.DB.Find(&fileFolder, "id = ?", fid)
	return fileFolder
}

// GetChildrenFolders 获取所有子文件夹
func GetChildrenFolders(parentId int, storeId int) (fileFolders []FileFolder) {
	mysql.DB.Order("time desc").Find(&fileFolders, "parent_folder_id = ? and file_store_id = ?", parentId, storeId)
	return fileFolders
}

// GetFolderParents 获取当前文件夹的所有父文件夹
func GetFolderParents(folder FileFolder) (fileFolders []FileFolder) {
	var cur = folder
	var par FileFolder
	for cur.ParentFolderId != 0 {
		mysql.DB.Find(&par, "id = ?", cur.ParentFolderId)
		fileFolders = append(fileFolders, par)
		cur = par
	}
	// 反转切片
	n := len(fileFolders)
	for i := 0; i < n/2; i++ {
		fileFolders[i], fileFolders[n-i-1] = fileFolders[n-i-1], fileFolders[i]
	}
	return fileFolders
}

// GetFolderPath 获取文件夹相对存储位置的路径
func GetFolderPath(folder FileFolder) string {
	path := "/" + strconv.Itoa(folder.FileStoreId) + "/"
	parents := GetFolderParents(folder)
	for _, value := range parents {
		path += value.FileFolderName + "/"
	}
	path += folder.FileFolderName
	return path
}

// GetUserFolderNum 获取用户仓库文件夹的数目
func GetUserFolderNum(fileStoreId int) (num int64) {
	var fileFolder []FileFolder
	mysql.DB.Find(&fileFolder, "file_store_id = ?", fileStoreId).Count(&num)
	return num
}

// DeleteFileFolder 删除文件夹（及子文件夹）
func DeleteFileFolder(fId int) {
	var fileFolders []FileFolder
	var files []MyFile
	//删除文件夹信息
	mysql.DB.Where("id = ?", fId).Delete(FileFolder{})
	//删除文件夹中文件信息
	mysql.DB.Find(&files, "parent_folder_id = ?", fId)
	for _, file := range files {
		DeleteFileById(file.Id)
	}
	//递归删除文件夹子文件夹信息
	mysql.DB.Find(&fileFolders, "parent_folder_id = ?", fId)
	for _, fileFolder := range fileFolders {
		DeleteFileFolder(fileFolder.Id)
	}
}

// DeleteStoreAllFolder 删除文件仓库中所有文件夹
func DeleteStoreAllFolder(fileStoreId int) {
	mysql.DB.Where("file_store_id = ?", fileStoreId).Delete(FileFolder{})
}

// UpdateFolderName 修改文件夹名
func UpdateFolderName(fId int, fName string) {
	var fileFolder FileFolder
	mysql.DB.Model(&fileFolder).Where("id = ?", fId).Update("file_folder_name", fName)
}

// FolderNameExists 判断父文件夹下改文件夹名是否存在
func FolderNameExists(parentId int, folderName string) bool {
	var folder FileFolder
	mysql.DB.Find(&folder, "file_folder_name = ? and parent_folder_id = ?", folderName, parentId)
	return folder.FileFolderName != ""
}
