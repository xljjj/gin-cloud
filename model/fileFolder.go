package model

// FileFolder 文件夹
type FileFolder struct {
	Id             int
	FileFolderName string
	ParentFolderId int
	FileStoreId    int
	Time           string
}
