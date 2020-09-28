package upload

import (
	"context"
	"errors"
	"mime/multipart"
	"server/library/global"

	"github.com/gogf/gf/frame/g"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct{}

// Upload 上传文件
func (*Minio) Upload(file *multipart.FileHeader) (string, string, error) {
	client, newErr := minio.New(global.Config.Minio.Endpoint, &minio.Options{Creds: credentials.NewStaticV4(global.Config.Minio.Id, global.Config.Minio.Secret, global.Config.Minio.Token), Secure: global.Config.Minio.UseSSL}) // Initialize minio client object.
	if newErr != nil {
		g.Log().Errorf("err:%v", newErr)
		return "", "", errors.New("function oss.New() Filed, err:" + newErr.Error())
	}
	ctx := context.Background()
	if bucketErr := client.MakeBucket(ctx, global.Config.Minio.Bucket, minio.MakeBucketOptions{Region: ""}); bucketErr != nil {
		if exists, existsErr := client.BucketExists(ctx, global.Config.Minio.Bucket); !exists && existsErr != nil { // Check to see if we already own this bucket (which happens if you run this twice)
			g.Log().Errorf("err:%v", existsErr)
			return "", "", errors.New("function client.BucketExists() Filed, err:" + existsErr.Error())
		}
		g.Log().Printf("We already own %s\n", global.Config.Minio.Bucket)
	} else {
		g.Log().Printf("Successfully created %s\n", global.Config.Minio.Bucket)
	}

	objectName := getObjectName(file.Filename)

	f, openErr := file.Open()
	if openErr != nil {
		g.Log().Errorf("err: %v", openErr.Error())
		return "", "", errors.New("function file.Open() Filed, err:" + openErr.Error())
	}

	// 获取文件类型
	contentType := file.Header.Get("content-type")

	info, putErr := client.PutObject(ctx, global.Config.Minio.Bucket, objectName, f, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if putErr != nil {
		g.Log().Errorf("err: %v", putErr.Error())
		return "", "", errors.New("function client.PutObject() Filed, err:" + putErr.Error())
	}

	g.Log().Printf("Successfully uploaded %s of size %d\n", objectName, info)
	return global.Config.Minio.Path + "/" + global.Config.Minio.Bucket + "/" + objectName, objectName, nil
}

// DeleteFile 删除文件
func (*Minio) DeleteFile(key string) error {
	var client *minio.Client
	client, newErr := minio.New(global.Config.Minio.Endpoint, &minio.Options{Creds: credentials.NewStaticV4(global.Config.Minio.Id, global.Config.Minio.Secret, global.Config.Minio.Token), Secure: global.Config.Minio.UseSSL}) // Initialize minio client object.
	if newErr != nil {
		g.Log().Errorf("err:%v", newErr)
		return errors.New("function oss.New() Filed, err:" + newErr.Error())
	}
	opts := minio.RemoveObjectOptions{GovernanceBypass: true}
	removeErr := client.RemoveObject(context.Background(), global.Config.Minio.Bucket, key, opts)
	if removeErr != nil {
		g.Log().Errorf("err: %v", removeErr.Error())
		return errors.New("function client.RemoveObject() Filed, err:" + removeErr.Error())
	}
	return nil
}
