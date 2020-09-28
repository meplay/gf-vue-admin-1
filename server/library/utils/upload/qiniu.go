package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var (
	qZone          string
	qBucket        string
	qImgPath       string
	qUseHTTPS      bool
	qAccessKey     string
	qSecretKey     string
	qUseCdnDomains bool
)

func init() {
	qZone = g.Cfg("oss").GetString("qiniu.Zone")
	qBucket = g.Cfg("oss").GetString("qiniu.Bucket")
	qImgPath = g.Cfg("oss").GetString("qiniu.ImgPath")
	qUseHTTPS = g.Cfg("oss").GetBool("qiniu.UseHTTPS")
	qAccessKey = g.Cfg("oss").GetString("qiniu.AccessKey")
	qSecretKey = g.Cfg("oss").GetString("qiniu.SecretKey")
	qUseCdnDomains = g.Cfg("oss").GetBool("qiniu.UseCdnDomains")
}

type Qiniu struct{}

// 接收两个参数 一个文件流 一个 bucket 你的七牛云标准空间的名字
func (*Qiniu) Upload(file *multipart.FileHeader) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: qBucket}
	mac := qbox.NewMac(qAccessKey, qSecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, openError := file.Open()
	if openError != nil {
		g.Log().Errorf("err:%v", openError)
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if putErr != nil {
		g.Log().Errorf("err:%v", putErr)
		return "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}
	return qImgPath + "/" + ret.Key, ret.Key, nil
}

func (*Qiniu) DeleteFile(key string) error {
	mac := qbox.NewMac(qAccessKey, qSecretKey)
	cfg := qiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	deleteErr := bucketManager.Delete(qAccessKey, key)
	if deleteErr != nil {
		g.Log().Errorf("err:%v", deleteErr)
		return errors.New("function file.Open() Filed, err:" + deleteErr.Error())
	}
	return nil
}

// config 根据配置文件进行返回七牛云的配置
func qiniuConfig() *storage.Config {
	cfg := storage.Config{UseHTTPS: qUseHTTPS, UseCdnDomains: qUseCdnDomains}
	switch qZone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
