package util

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"io"
	"path"
)

// UploadOss 上传文件至阿里云
func UploadOss(filename, fileHash string) {
	//获取文件后缀
	fileSuffix := path.Ext(filename)
	// 创建OSSClient实例
	client, err := oss.New(viper.GetString("oss.endPoint"),
		viper.GetString("oss.accessKeyId"), viper.GetString("oss.accessKeySecret"))
	if err != nil {
		fmt.Println("CreateClient Error:", err)
		return
	}

	// 获取存储空间
	bucket, err := client.Bucket(viper.GetString("oss.bucketName"))
	if err != nil {
		fmt.Println("Get Space Error:", err)
		return
	}

	// 上传本地文件
	err = bucket.PutObjectFromFile("file/"+fileHash+fileSuffix,
		viper.GetString("upload.location")+filename)
	if err != nil {
		fmt.Println("Local Upload Error:", err)
		return
	}
}

// DownloadOss 从oss下载文件
func DownloadOss(fileHash, fileType string) []byte {
	// 创建OSSClient实例
	client, err := oss.New(viper.GetString("oss.endPoint"),
		viper.GetString("oss.accessKeyId"), viper.GetString("oss.accessKeySecret"))
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 获取存储空间
	bucket, err := client.Bucket(viper.GetString("oss.bucketName"))
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 下载文件到流
	body, err := bucket.GetObject("file/" + fileHash + fileType)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return data
}

// DeleteOss 从oss删除文件
func DeleteOss(fileHash, fileType string) {
	// 创建OSSClient实例
	client, err := oss.New(viper.GetString("oss.endPoint"),
		viper.GetString("oss.accessKeyId"), viper.GetString("oss.accessKeySecret"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 获取存储空间
	bucket, err := client.Bucket(viper.GetString("oss.bucketName"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = bucket.DeleteObject("file/" + fileHash + fileType)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
