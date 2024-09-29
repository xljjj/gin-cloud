package util

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"io"
	"path/filepath"
)

// UploadOss 上传文件至阿里云
func UploadOss(filePath, fileHash, Suffix string) error {
	// 创建OSSClient实例
	client, err := oss.New(viper.GetString("oss.endPoint"),
		viper.GetString("oss.accessKeyId"), viper.GetString("oss.accessKeySecret"))
	if err != nil {
		return err
	}

	// 获取存储空间
	bucket, err := client.Bucket(viper.GetString("oss.bucketName"))
	if err != nil {
		return err
	}

	// 上传本地文件
	err = bucket.PutObjectFromFile("file/"+fileHash+Suffix,
		filepath.Join(viper.GetString("oss.local"), filePath))
	if err != nil {
		return err
	}
	return nil
}

// DownloadOss 从oss下载文件
func DownloadOss(fileHash, Suffix string) ([]byte, error) {
	// 创建OSSClient实例
	client, err := oss.New(viper.GetString("oss.endPoint"),
		viper.GetString("oss.accessKeyId"), viper.GetString("oss.accessKeySecret"))
	if err != nil {
		return nil, err
	}

	// 获取存储空间
	bucket, err := client.Bucket(viper.GetString("oss.bucketName"))
	if err != nil {
		return nil, err
	}

	// 下载文件到流
	body, err := bucket.GetObject("file/" + fileHash + Suffix)
	if err != nil {
		return nil, err
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// DeleteOss 从oss删除文件
func DeleteOss(fileHash, Suffix string) error {
	// 创建OSSClient实例
	client, err := oss.New(viper.GetString("oss.endPoint"),
		viper.GetString("oss.accessKeyId"), viper.GetString("oss.accessKeySecret"))
	if err != nil {
		return err
	}

	// 获取存储空间
	bucket, err := client.Bucket(viper.GetString("oss.bucketName"))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	err = bucket.DeleteObject("file/" + fileHash + Suffix)
	if err != nil {
		return err
	}
	return nil
}
