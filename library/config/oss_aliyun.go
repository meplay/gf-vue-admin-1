package config

import "time"

type Aliyun struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	BasePath        string `mapstructure:"base-path" json:"basePath" yaml:"base-path"`
	BucketUrl       string `mapstructure:"bucket-url" json:"bucketUrl" yaml:"bucket-url"`
	BucketName      string `mapstructure:"bucket-name" json:"bucketName" yaml:"bucket-name"`
	AccessKeyId     string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"accessKeySecret" yaml:"access-key-secret"`
}

// Filepath 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *Aliyun) Filepath(filename string) string {
	return a.BasePath + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + filename
}