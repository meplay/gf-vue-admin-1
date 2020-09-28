package upload

import (
	"fmt"
	"mime/multipart"
	"server/library/global"
	"time"
)

var Oss OSS

type OSS interface {
	Upload(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

func getObjectName(filename string) string {
	folder := time.Now().Format("20060102")
	return fmt.Sprintf("%s/%d%s", folder, time.Now().Unix(), filename) // 文件名格式 自己可以改 建议保证唯一性
}

func InitUpload() {
	switch global.Config.System.OssType {
	case "local":
		Oss = &Local{}
	case "qiniu":
		Oss = &Qiniu{}
	case "minio":
		Oss = &Minio{}
	case "aliyun":
		Oss = &AliYun{}
	default:
		Oss = &Local{}
	}
}
