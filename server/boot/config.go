package boot

import (
	"gf-vue-admin/library/config"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
)

var Config = new(_config)

type _config struct {
	config.Config
}

func (c *_config) Initialize() {
	// Jwt
	c.Jwt.ExpiresAt = g.Cfg("jwt").GetInt("jwt.ExpiresAt")
	c.Jwt.RefreshAt = g.Cfg("jwt").GetInt("jwt.RefreshAt")
	c.Jwt.SigningKey = g.Cfg("jwt").GetString("jwt.SigningKey")

	// Local
	c.Oss.Local.LocalPath = g.Cfg("oss").GetString("local.LocalPath")

	// Qiniu
	c.Oss.Qiniu.Zone = g.Cfg("oss").GetString("qiniu.Zone")
	c.Oss.Qiniu.Bucket = g.Cfg("oss").GetString("qiniu.Bucket")
	c.Oss.Qiniu.ImgPath = g.Cfg("oss").GetString("qiniu.ImgPath")
	c.Oss.Qiniu.UseHTTPS = g.Cfg("oss").GetBool("qiniu.UseHTTPS")
	c.Oss.Qiniu.AccessKey = g.Cfg("oss").GetString("qiniu.AccessKey")
	c.Oss.Qiniu.SecretKey = g.Cfg("oss").GetString("qiniu.SecretKey")
	c.Oss.Qiniu.UseCdnDomains = g.Cfg("oss").GetBool("qiniu.UseCdnDomains")

	// Minio
	c.Oss.Minio.Id = g.Cfg("oss").GetString("minio.Id")
	c.Oss.Minio.Path = g.Cfg("oss").GetString("minio.Path")
	c.Oss.Minio.Token = g.Cfg("oss").GetString("minio.Token")
	c.Oss.Minio.Bucket = g.Cfg("oss").GetString("minio.Bucket")
	c.Oss.Minio.UseSsl = g.Cfg("oss").GetBool("minio.UseSsl")
	c.Oss.Minio.Secret = g.Cfg("oss").GetString("minio.Secret")
	c.Oss.Minio.Endpoint = g.Cfg("oss").GetString("minio.Endpoint")

	// Aliyun
	c.Oss.Aliyun.Path = g.Cfg("oss").GetString("aliyun.Path")
	c.Oss.Aliyun.Bucket = g.Cfg("oss").GetString("aliyun.Bucket")
	c.Oss.Aliyun.ACLType = g.Cfg("oss").GetString("aliyun.ACLType")
	c.Oss.Aliyun.Endpoint = g.Cfg("oss").GetString("aliyun.Endpoint")
	c.Oss.Aliyun.AccessKeyID = g.Cfg("oss").GetString("aliyun.AccessKeyID")
	c.Oss.Aliyun.SecretAccessKey = g.Cfg("oss").GetString("aliyun.SecretAccessKey")
	c.Oss.Aliyun.StorageClassType = g.Cfg("oss").GetString("aliyun.StorageClassType")

	// Redis
	c.Redis.Default = g.Cfg().GetString("redis.default")
	c.Redis.Cache = g.Cfg().GetString("redis.cache")

	// Email
	c.Email.To = g.Cfg("email").GetString("email.To")
	c.Email.Port = g.Cfg("email").GetInt("email.Port")
	c.Email.From = g.Cfg("email").GetString("email.From")
	c.Email.Host = g.Cfg("email").GetString("email.Host")
	c.Email.IsSsl = g.Cfg("email").GetBool("email.IsSsl")
	c.Email.Secret = g.Cfg("email").GetString("email.Secret")
	c.Email.Nickname = g.Cfg("email").GetString("email.Nickname")

	// Casbin
	c.Casbin.ModelPath = g.Cfg("casbin").GetString("casbin.ModelPath")

	// Logger
	c.Logger.Path = g.Cfg().GetString("logger.Path")
	c.Logger.Level = g.Cfg().GetString("logger.Level")
	c.Logger.Stdout = g.Cfg().GetBool("logger.Stdout")

	// System
	c.System.Db = g.Cfg("system").GetString("system.Db")
	c.System.Env = g.Cfg("system").GetString("system.Env")
	c.System.OssType = g.Cfg("system").GetString("system.OssType")
	c.System.ErrorToEmail = g.Cfg("system").GetBool("system.ErrorToEmail")
	c.System.UseMultipoint = g.Cfg("system").GetBool("system.UseMultipoint")

	// Server
	c.Server.LogPath = g.Cfg().GetString("server.LogPath")
	c.Server.Address = g.Cfg().GetString("server.Address")
	c.Server.DumpRouterMap = g.Cfg().GetBool("server.DumpRouterMap")
	c.Server.ErrorLogEnabled = g.Cfg().GetBool("server.ErrorLogEnabled")
	c.Server.AccessLogEnabled = g.Cfg().GetBool("server.AccessLogEnabled")

	// Captcha
	c.Captcha.KeyLong = g.Cfg("captcha").GetInt("captcha.KeyLong")
	c.Captcha.ImageWidth = g.Cfg("captcha").GetInt("captcha.ImageWidth")
	c.Captcha.ImageHeight = g.Cfg("captcha").GetInt("captcha.ImageHeight")

	// Database
	c.Database.Host = g.Cfg().GetString("database.default.Host")
	c.Database.Port = g.Cfg().GetString("database.default.Port")
	c.Database.User = g.Cfg().GetString("database.default.User")
	c.Database.Pass = g.Cfg().GetString("database.default.Pass")
	c.Database.Name = g.Cfg().GetString("database.default.Name")
	c.Database.Type = g.Cfg().GetString("database.default.Type")
	c.Database.Role = g.Cfg().GetString("database.default.Role")
	c.Database.Debug = g.Cfg().GetBool("database.default.Debug")
	c.Database.Prefix = g.Cfg().GetString("database.default.Prefix")
	c.Database.DryRun = g.Cfg().GetBool("database.default.DryRun")
	c.Database.Weight = g.Cfg().GetInt("database.default.Weight")
	c.Database.Charset = g.Cfg().GetString("database.default.Charset")
	c.Database.LinkInfo = g.Cfg().GetString("database.default.LinkInfo")
	c.Database.MaxIdleConnCount = g.Cfg().GetInt("database.default.MaxIdleConnCount")
	c.Database.MaxOpenConnCount = g.Cfg().GetInt("database.default.MaxOpenConnCount")
	c.Database.MaxConnLifetime = g.Cfg().GetDuration("database.default.MaxConnLifetime")

	// DatabaseLogger
	c.DatabaseLogger.Path = g.Cfg().GetString("database.logger.Path")
	c.DatabaseLogger.Level = g.Cfg().GetString("database.logger.Level")
	c.DatabaseLogger.Stdout = g.Cfg().GetBool("database.logger.Stdout")
	global.Config = &c.Config
}
