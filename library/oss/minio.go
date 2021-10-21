package oss

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

var Minio = new(_minio)

type _minio struct {
	filesize    int64
	filename    string
	contentType string
}

func NewMinioClient() (client *minio.Client, err error) {
	client, err = minio.New(global.Config.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(global.Config.Minio.Id, global.Config.Minio.Secret, global.Config.Minio.Token),
		Secure: global.Config.Minio.UseSsl,
	}) // Initialize minio client object.
	if err != nil {
		return nil, errors.Wrap(err, "初始化minio Client 对象失败!")
	}
	return client, nil
}

func (m *_minio) DeleteByKey(key string) error {
	client, err := NewMinioClient()
	if err != nil {
		return err
	}
	options := minio.RemoveObjectOptions{GovernanceBypass: true}
	err = client.RemoveObject(context.Background(), global.Config.Minio.Bucket, key, options)
	if err != nil {
		return errors.Wrap(err, "删除文件失败!")
	}
	return nil
}

func (m *_minio) UploadByFile(file multipart.File) (filepath string, filename string, err error) {
	var client *minio.Client
	if client, err = NewMinioClient(); err != nil {
		return filepath, filename, err
	}

	ctx := context.Background()
	if err = client.MakeBucket(ctx, global.Config.Minio.Bucket, minio.MakeBucketOptions{Region: ""}); err != nil {
		if exists, existsErr := client.BucketExists(ctx, global.Config.Minio.Bucket); !exists && existsErr != nil { // Check to see if we already own this bucket (which happens if you run this twice)
			g.Log().Error("function client.BucketExists() Failed!", g.Map{"err": err})
			return filepath, filename, err
		}
		zap.L().Info("We Already Own!", zap.String("bucket", global.Config.Minio.Bucket))
	} else {
		zap.L().Info("Successfully Created!", zap.String("bucket", global.Config.Minio.Bucket))
	}

	filename = global.Config.Minio.Filename(m.filename)
	filepath = global.Config.Minio.Filepath(filename)

	var info minio.UploadInfo
	if info, err = client.PutObject(ctx, global.Config.Minio.Bucket, filename, file, m.filesize, minio.PutObjectOptions{ContentType: m.contentType}); err != nil {
		return filepath, filename, errors.Wrap(err, "上传minio服务器失败!")
	}
	zap.L().Info("Successfully uploaded!", zap.String("filepath", filepath), zap.String("filename", filename), zap.Int64("filesize", info.Size))
	return filepath, filename, nil
}

func (m *_minio) UploadByFilepath(p string) (path string, filename string, err error) {
	var file *os.File
	file, err = os.Open(p)
	if err != nil {
		return path, filename, errors.Wrapf(err, "(%s)文件不存在!", p)
	}
	var info os.FileInfo
	info, err = file.Stat()
	if err != nil {
		return path, filename, errors.Wrapf(err, "(%s)文件信息获取失败!", p)
	}
	m.filesize = info.Size()
	_, m.filename = filepath.Split(path)
	return m.UploadByFile(file)
}

func (m *_minio) UploadByFileHeader(header *multipart.FileHeader) (filepath string, filename string, err error) {
	var open multipart.File
	open, err = header.Open()
	if err != nil {
		return filepath, filename, err
	}
	m.contentType = header.Header.Get("content-type") // 获取文件类型
	m.filename = fmt.Sprintf("%d%s", time.Now().Unix(), header.Filename)
	m.filesize = header.Size
	return m.UploadByFile(open)
}
