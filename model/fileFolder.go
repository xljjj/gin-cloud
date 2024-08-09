package model

import (
	"CloudDrive/util"
	"fmt"
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

// createFileFolder 新建文件夹
func createFileFolder(folderName string, parentId string, fileStoreId int) {
	parentIntId, err := strconv.Atoi(parentId)
	if err != nil {
		fmt.Println("父文件夹ID非法")
		return
	}
	fileFolder := FileFolder{
		FileFolderName: folderName,
		ParentFolderId: parentIntId,
		FileStoreId:    fileStoreId,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
	}
	util.DB.Create(&fileFolder)
}

// GetFolderById 根据ID得到文件夹
func GetFolderById(fid int) (fileFolder FileFolder) {
	util.DB.Find(&fileFolder, "id = ?", fid)
	return fileFolder
}

// GetChildrenFolders 获取所有子文件夹
func GetChildrenFolders(parentId int, storeId int) (fileFolders []FileFolder) {
	util.DB.Order("time desc").Find(&fileFolders, "parent_folder_id = ? and file_store_id = ?", parentId, storeId)
	return fileFolders
}

// GetPathParents 获取当前路径所有父文件夹
func GetPathParents(folder FileFolder) (fileFolders []FileFolder) {
	var cur = folder
	var par FileFolder
	for cur.ParentFolderId != 0 {
		util.DB.Find(&par, "id = ?", cur.ParentFolderId)
		fileFolders = append(fileFolders, par)
		cur = par
	}
	//反转
	for i, j := 0, len(fileFolders)-1; i < j; i, j = i+1, j-1 {
		fileFolders[i], fileFolders[j] = fileFolders[j], fileFolders[i]
	}
	return fileFolders
}

// GetUserFolderNum 获取用户仓库文件夹的数目
func GetUserFolderNum(fileStoreId int) (num int64) {
	var fileFolder []FileFolder
	util.DB.Find(&fileFolder, "file_store_id = ?", fileStoreId).Count(&num)
	return num
}

// DeleteFileFolder 删除文件夹（及子文件夹）
func DeleteFileFolder(fId int) {
	var fileFolder FileFolder
	var fileFolder2 FileFolder
	//删除文件夹信息
	util.DB.Where("id = ?", fId).Delete(FileFolder{})
	//删除文件夹中文件信息
	util.DB.Where("parent_folder_id = ?", fId).Delete(MyFile{})
	//删除文件夹中文件夹信息
	util.DB.Find(&fileFolder, "parent_folder_id = ?", fId)
	util.DB.Where("parent_folder_id = ?", fId).Delete(FileFolder{})

	util.DB.Find(&fileFolder2, "parent_folder_id = ?", fileFolder.Id)
	//递归删除文件下的文件夹
	if fileFolder2.Id != 0 {
		DeleteFileFolder(fileFolder.Id)
	}
}

// UpdateFolderName 修改文件夹名
func UpdateFolderName(fId int, fName string) {
	var fileFolder FileFolder
	util.DB.Model(&fileFolder).Where("id = ?", fId).Update("file_folder_name", fName)
}
