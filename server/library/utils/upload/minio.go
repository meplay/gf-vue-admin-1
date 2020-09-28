package upload

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/gogf/gf/frame/g"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	mId       string
	mPath     string
	mToken    string
	mBucket   string
	mSecret   string
	mEndpoint string
	mUseSSL   bool
)

func init() {
	mId = g.Cfg("oss").GetString("minio.Id")
	mPath = g.Cfg("oss").GetString("minio.Path")
	mToken = g.Cfg("oss").GetString("minio.Token")
	mBucket = g.Cfg("oss").GetString("minio.Bucket")
	mUseSSL = g.Cfg("oss").GetBool("minio.UseSsl")
	mSecret = g.Cfg("oss").GetString("minio.Secret")
	mEndpoint = g.Cfg("oss").GetString("minio.Endpoint")
}

type Minio struct{}

// Upload 上传文件
func (*Minio) Upload(file *multipart.FileHeader) (string, string, error) {
	client, newErr := minio.New(mEndpoint, &minio.Options{Creds: credentials.NewStaticV4(mId, mSecret, mToken), Secure: mUseSSL}) // Initialize minio client object.
	if newErr != nil {
		g.Log().Errorf("err:%v", newErr)
		return "", "", errors.New("function oss.New() Filed, err:" + newErr.Error())
	}
	ctx := context.Background()
	if bucketErr := client.MakeBucket(ctx, mBucket, minio.MakeBucketOptions{Region: ""}); bucketErr != nil {
		if exists, existsErr := client.BucketExists(ctx, mBucket); !exists && existsErr != nil { // Check to see if we already own this bucket (which happens if you run this twice)
			g.Log().Errorf("err:%v", existsErr)
			return "", "", errors.New("function client.BucketExists() Filed, err:" + existsErr.Error())
		}
		g.Log().Printf("We already own %s\n", mBucket)
	} else {
		g.Log().Printf("Successfully created %s\n", mBucket)
	}

	objectName := getObjectName(file.Filename)

	f, openErr := file.Open()
	if openErr != nil {
		g.Log().Errorf("err: %v", openErr.Error())
		return "", "", errors.New("function file.Open() Filed, err:" + openErr.Error())
	}

	// 获取文件类型
	contentType := file.Header.Get("content-type")

	info, putErr := client.PutObject(ctx, mBucket, objectName, f, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if putErr != nil {
		g.Log().Errorf("err: %v", putErr.Error())
		return "", "", errors.New("function client.PutObject() Filed, err:" + putErr.Error())
	}

	g.Log().Printf("Successfully uploaded %s of size %d\n", objectName, info)
	return mPath + "/" + mBucket + "/" + objectName, objectName, nil
}

// DeleteFile 删除文件
func (*Minio) DeleteFile(key string) error {
	var client *minio.Client
	client, newErr := minio.New(mEndpoint, &minio.Options{Creds: credentials.NewStaticV4(mId, mSecret, mToken), Secure: mUseSSL}) // Initialize minio client object.
	if newErr != nil {
		g.Log().Errorf("err:%v", newErr)
		return errors.New("function oss.New() Filed, err:" + newErr.Error())
	}
	opts := minio.RemoveObjectOptions{GovernanceBypass: true}
	removeErr := client.RemoveObject(context.Background(), mBucket, key, opts)
	if removeErr != nil {
		g.Log().Errorf("err: %v", removeErr.Error())
		return errors.New("function client.RemoveObject() Filed, err:" + removeErr.Error())
	}
	return nil
}
