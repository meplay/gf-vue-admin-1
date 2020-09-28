package global

import (
	"server/config"

	"github.com/gogf/gf/frame/g"
)

var Config *config.Server

func init() {
	minio()
	qiniu()
	aliyun()
}

func minio() {
	Config.Minio.Id = g.Cfg().GetString("minio.Id")
	Config.Minio.Path = g.Cfg().GetString("minio.Path")
	Config.Minio.Token = g.Cfg().GetString("minio.Token")
	Config.Minio.Bucket = g.Cfg().GetString("minio.Bucket")
	Config.Minio.UseSSL = g.Cfg().GetBool("minio.UseSsl")
	Config.Minio.Secret = g.Cfg().GetString("minio.Secret")
	Config.Minio.Endpoint = g.Cfg().GetString("minio.Endpoint")
}

func qiniu() {
	Config.Qiniu.Zone = g.Cfg().GetString("qiniu.Zone")
	Config.Qiniu.Bucket = g.Cfg().GetString("qiniu.Bucket")
	Config.Qiniu.ImgPath = g.Cfg().GetString("qiniu.ImgPath")
	Config.Qiniu.UseHTTPS = g.Cfg().GetBool("qiniu.UseHTTPS")
	Config.Qiniu.AccessKey = g.Cfg().GetString("qiniu.AccessKey")
	Config.Qiniu.SecretKey = g.Cfg().GetString("qiniu.SecretKey")
	Config.Qiniu.UseCdnDomains = g.Cfg().GetBool("qiniu.UseCdnDomains")
}

func aliyun() {
	Config.AliYun.Path = g.Cfg().GetString("aliyun.Path")
	Config.AliYun.Bucket = g.Cfg().GetString("aliyun.Bucket")
	Config.AliYun.Endpoint = g.Cfg().GetString("aliyun.Endpoint")
	Config.AliYun.AccessKeyID = g.Cfg().GetString("aliyun.AccessKeyID")
	Config.AliYun.SecretAccessKey = g.Cfg().GetString("aliyun.SecretAccessKey")
	Config.AliYun.StorageClassType = g.Cfg().GetString("aliyun.StorageClassType")
}
