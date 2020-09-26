package utils

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gogf/gf/frame/g"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var c *config

type config struct {
	Id       string
	Path     string
	Token    string
	Bucket   string
	Secret   string
	Endpoint string
	UseSSL   bool
}

func initConfig() {
	c = (*config)(nil)
	c.Id = g.Cfg().GetString("minio.Id")
	c.Path = g.Cfg().GetString("minio.Path")
	c.Token = g.Cfg().GetString("minio.Token")
	c.Bucket = g.Cfg().GetString("minio.Bucket")
	c.Secret = g.Cfg().GetString("minio.Secret")
	c.Endpoint = g.Cfg().GetString("minio.Endpoint")
	c.UseSSL = g.Cfg().GetBool("minio.UseSsl")
}

type Minio struct{}

func minioInit() (client *minio.Client, err error) {
	client, err = minio.New(c.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.Id, c.Secret, c.Token),
		Secure: c.UseSSL,
	})
	if err != nil {

	}
	return
}

// 接收两个参数 一个文件流 一个 bucket 你的七牛云标准空间的名字
func (oss *Minio) Upload(file *multipart.FileHeader) (path string, key string, err error) {
	var (
		exists bool
		f      multipart.File
		info   minio.UploadInfo
		client *minio.Client
	)
	if client, err = minioInit(); err != nil { // Initialize minio client object.
		g.Log().Error(err)
		return "", "", errors.New("建立client失败")
	}
	ctx := context.Background()
	if err = client.MakeBucket(ctx, c.Bucket, minio.MakeBucketOptions{Region: ""}); err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		if exists, err = client.BucketExists(ctx, c.Bucket); !exists && err != nil {
			g.Log().Errorf("err: %v", err.Error())
			return "", "", errors.New("client.BucketExists 失败")
		}
		g.Log().Printf("We already own %s\n", c.Bucket)
	} else {
		g.Log().Printf("Successfully created %s\n", c.Bucket)
	}

	name := fmt.Sprintf("%s/%d%s", time.Now().Format("20060102"), time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	if f, err = file.Open(); err != nil {
		g.Log().Errorf("err: %v", err.Error())
		return "", "", errors.New("文件Open失败")
	}

	// 获取文件类型
	contentType := file.Header.Get("content-type")

	if info, err = client.PutObject(ctx, c.Bucket, name, f, file.Size, minio.PutObjectOptions{ContentType: contentType}); err != nil {
		g.Log().Errorf("err: %v", err.Error())
		return "", "", errors.New("client.PutObject 失败")
	}

	g.Log().Printf("Successfully uploaded %s of size %d\n", name, info)
	return c.Path + "/" + c.Bucket + "/" + name, name, err
}

func (oss *Minio) DeleteFile(key string) (err error) {
	var client *minio.Client
	if client, err = minioInit(); err != nil { // Initialize minio client object.
		g.Log().Error(err)
		return errors.New("建立client失败")
	}
	opts := minio.RemoveObjectOptions{GovernanceBypass: true}
	err = client.RemoveObject(context.Background(), c.Bucket, key, opts)
	if err != nil {
		g.Log().Errorf("err: %v", err.Error())
		return errors.New("client.RemoveObject 失败")
	}
	return
}
