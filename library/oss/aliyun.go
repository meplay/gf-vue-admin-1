package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"mime/multipart"
)

var _ interfaces.Oss = (*aliyun)(nil)

var Aliyun = new(aliyun)

type aliyun struct {
	filename string
}

func NewAliyunBucket() (bucket *oss.Bucket, err error) {
	var client *oss.Client
	if client, err = oss.New(global.Config.Aliyun.Endpoint, global.Config.Aliyun.AccessKeyId, global.Config.Aliyun.AccessKeySecret); err != nil {
		return nil, err
	} // 创建OSSClient实例

	if bucket, err = client.Bucket(global.Config.Aliyun.BucketName); err != nil {
		return nil, errors.Wrap(err, "获取存储空间失败!")
	} // 获取存储空间

	return bucket, nil
}

func (a *aliyun) DeleteByKey(key string) error {
	bucket, err := NewAliyunBucket()
	if err != nil {
		return err
	}

	if err = bucket.DeleteObject(key); err != nil {
		return errors.Wrap(err, "删除文件失败!")
	}

	return nil
}

func (a *aliyun) UploadByFile(file multipart.File) (filepath string, filename string, err error) {
	bucket, newErr := NewAliyunBucket()
	if newErr != nil {
		return filepath, filename, newErr
	}

	defer func() {
		if err = file.Close(); err != nil {
			zap.L().Error("文件关闭失败!", zap.Error(err))
		}
	}() // 关闭文件流

	filepath = global.Config.Aliyun.Filepath(a.filename)

	err = bucket.PutObject(filepath, file) // 上传文件流。
	if err != nil {
		return filepath, filename, errors.Wrap(err, "上传文件流失败!")
	}

	return global.Config.Aliyun.BucketUrl + "/" + filepath, a.filename, nil
}

func (a *aliyun) UploadByFileHeader(file *multipart.FileHeader) (filepath string, filename string, err error) {
	var open multipart.File
	open, err = file.Open()
	if err != nil {
		return filepath, filename, err
	}
	a.filename = file.Filename
	return a.UploadByFile(open)
}
