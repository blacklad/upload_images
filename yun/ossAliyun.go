package yun

import (
	"blacklad.com/upload_images/conf"
	"blacklad.com/upload_images/utils"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"path/filepath"
)

const OssDirSeparator = "/"

type OssAli struct {
	basePath     string
	BucketClient *oss.Bucket
}

func NewOssAli(config *conf.Config) (*OssAli, error) {
	// 创建OSSClient实例。
	client, err := oss.New(config.OssConfig.Endpoint, config.OssConfig.Key, config.OssConfig.Secret)
	if err != nil {
		utils.HandleError(err)
	}
	// 获取存储空间。
	bucketName := "blacklad"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return &OssAli{
		basePath:     config.OssBasePath,
		BucketClient: bucket,
	}, nil
}

// 上传文件到oss。
func (o *OssAli) UploadFile(filePath, localPath string) error {
	ossPath := filepath.Join(o.basePath, filePath)
	// 创建文件元信息。
	options := []oss.Option{
		oss.ObjectACL(oss.ACLPublicRead),
	}
	err := o.BucketClient.PutObjectFromFile(ossPath, localPath, options...)
	return err
}
