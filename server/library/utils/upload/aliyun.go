package upload

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/gogf/gf/frame/g"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	aPath             string
	aBucket           string
	aEndpoint         string
	aAccessKeyID      string
	aSecretAccessKey  string
	aStorageClassType string
)

func init() {
	aPath = g.Cfg("oss").GetString("aliyun.Path")
	aBucket = g.Cfg("oss").GetString("aliyun.Bucket")
	aEndpoint = g.Cfg("oss").GetString("aliyun.Endpoint")
	aAccessKeyID = g.Cfg("oss").GetString("aliyun.AccessKeyID")
	aSecretAccessKey = g.Cfg("oss").GetString("aliyun.SecretAccessKey")
	aStorageClassType = g.Cfg("oss").GetString("aliyun.StorageClassType")
}

type AliYun struct{}

func (*AliYun) Upload(file *multipart.FileHeader) (string, string, error) {
	var storageType oss.Option
	client, newErr := oss.New(aEndpoint, aAccessKeyID, aSecretAccessKey, oss.Timeout(10, 120))
	if newErr != nil {
		g.Log().Errorf("err:%v", newErr)
		return "", "", errors.New("function oss.New() Filed, err:" + newErr.Error())
	}

	// 获取存储空间。
	bucket, bucketErr := client.Bucket(aBucket)
	if bucketErr != nil {
		g.Log().Errorf("err:%v", bucketErr)
		return "", "", errors.New("function client.Bucket() Filed, err:" + bucketErr.Error())
	}

	switch aStorageClassType { // 根据配置文件进行指定存储类型
	case "Standard": // 指定存储类型为标准存储，缺省也为标准存储。
		storageType = oss.ObjectStorageClass(oss.StorageStandard)
	case "IA": // 指定存储类型为很少访问存储
		storageType = oss.ObjectStorageClass(oss.StorageIA)
	case "Archive": // 指定存储类型为归档存储。
		storageType = oss.ObjectStorageClass(oss.StorageArchive)
	case "ColdArchive": // 指定存储类型为归档存储。
		storageType = oss.ObjectStorageClass(oss.StorageColdArchive)
	}

	f, openError := file.Open() // 读取文件
	if openError != nil {
		g.Log().Errorf("err:%v", openError)
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}

	// 获取文件类型
	contentType := file.Header.Get("content-type")
	objectType := oss.ContentType(contentType)

	// 指定访问权限为公共读，缺省为继承bucket的权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	// 文件对象名
	objectName := getObjectName(file.Filename)

	// 上传
	putErr := bucket.PutObject(objectName, f, storageType, objectType, objectAcl)
	if putErr != nil {
		g.Log().Errorf("err:%v", putErr)
		return "", "", errors.New("function bucket.PutObject() Filed, err:" + putErr.Error())
	}

	return aPath + "/" + objectName, objectName, nil
}

func (*AliYun) DeleteFile(key string) error {
	client, newErr := oss.New(aEndpoint, aAccessKeyID, aSecretAccessKey, oss.Timeout(10, 120))
	if newErr != nil {
		g.Log().Errorf("err:%v", newErr)
		return errors.New("function oss.New() Filed, err:" + newErr.Error())
	}

	// 获取存储空间。
	bucket, err := client.Bucket(aBucket)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		g.Log().Errorf("Delete File Failed, err:%v", newErr)
		return errors.New("function bucket.DeleteObject() Filed, err:" + newErr.Error())
	}
	return nil
}
