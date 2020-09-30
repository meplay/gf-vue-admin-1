package model

import "github.com/gogf/gf/frame/g"

var Config ServerConfig

func init() {
	Config.Jwt.ExpiresAt = g.Cfg("jwt").GetInt("jwt.ExpiresAt")
	Config.Jwt.RefreshAt = g.Cfg("jwt").GetInt("jwt.RefreshAt")
	Config.Jwt.SigningKey = g.Cfg("jwt").GetString("jwt.SigningKey")
}

func init() {
	Config.Oss.Local.LocalPath = g.Cfg("oss").GetString("local.LocalPath")

	Config.Oss.Qiniu.Zone = g.Cfg("oss").GetString("qiniu.Zone")
	Config.Oss.Qiniu.Bucket = g.Cfg("oss").GetString("qiniu.Bucket")
	Config.Oss.Qiniu.ImgPath = g.Cfg("oss").GetString("qiniu.ImgPath")
	Config.Oss.Qiniu.UseHTTPS = g.Cfg("oss").GetBool("qiniu.UseHTTPS")
	Config.Oss.Qiniu.AccessKey = g.Cfg("oss").GetString("qiniu.AccessKey")
	Config.Oss.Qiniu.SecretKey = g.Cfg("oss").GetString("qiniu.SecretKey")
	Config.Oss.Qiniu.UseCdnDomains = g.Cfg("oss").GetBool("qiniu.UseCdnDomains")

	Config.Oss.Minio.Id = g.Cfg("oss").GetString("minio.Id")
	Config.Oss.Minio.Path = g.Cfg("oss").GetString("minio.Path")
	Config.Oss.Minio.Token = g.Cfg("oss").GetString("minio.Token")
	Config.Oss.Minio.Bucket = g.Cfg("oss").GetString("minio.Bucket")
	Config.Oss.Minio.UseSsl = g.Cfg("oss").GetBool("minio.UseSsl")
	Config.Oss.Minio.Secret = g.Cfg("oss").GetString("minio.Secret")
	Config.Oss.Minio.Endpoint = g.Cfg("oss").GetString("minio.Endpoint")

	Config.Oss.Aliyun.Path = g.Cfg("oss").GetString("aliyun.Path")
	Config.Oss.Aliyun.Bucket = g.Cfg("oss").GetString("aliyun.Bucket")
	Config.Oss.Aliyun.ACLType = g.Cfg("oss").GetString("aliyun.ACLType")
	Config.Oss.Aliyun.Endpoint = g.Cfg("oss").GetString("aliyun.Endpoint")
	Config.Oss.Aliyun.AccessKeyID = g.Cfg("oss").GetString("aliyun.AccessKeyID")
	Config.Oss.Aliyun.SecretAccessKey = g.Cfg("oss").GetString("aliyun.SecretAccessKey")
	Config.Oss.Aliyun.StorageClassType = g.Cfg("oss").GetString("aliyun.StorageClassType")
}

func init() {
	Config.Redis.Default = g.Cfg().GetString("redis.default")
	Config.Redis.Cache = g.Cfg().GetString("redis.cache")
}

func init() {
	Config.Email.To = g.Cfg("email").GetString("email.To")
	Config.Email.Port = g.Cfg("email").GetInt("email.Port")
	Config.Email.From = g.Cfg("email").GetString("email.From")
	Config.Email.Host = g.Cfg("email").GetString("email.Host")
	Config.Email.IsSsl = g.Cfg("email").GetBool("email.IsSsl")
	Config.Email.Secret = g.Cfg("email").GetString("email.Secret")
	Config.Email.Nickname = g.Cfg("email").GetString("email.Nickname")
}

func init() {
	Config.Casbin.ModelPath = g.Cfg("casbin").GetString("casbin.ModelPath")
}

func init() {
	Config.Logger.Path = g.Cfg().GetString("logger.Path")
	Config.Logger.Level = g.Cfg().GetString("logger.Level")
	Config.Logger.Stdout = g.Cfg().GetBool("logger.Stdout")
}

func init() {
	Config.System.Db = g.Cfg("system").GetString("system.Db")
	Config.System.Env = g.Cfg("system").GetString("system.Env")
	Config.System.OssType = g.Cfg("system").GetString("system.OssType")
	Config.System.ErrorToEmail = g.Cfg("system").GetBool("system.ErrorToEmail")
	Config.System.UseMultipoint = g.Cfg("system").GetBool("system.UseMultipoint")
}

func init() {
	Config.Server.LogPath = g.Cfg().GetString("server.LogPath")
	Config.Server.Address = g.Cfg().GetString("server.Address")
	Config.Server.DumpRouterMap = g.Cfg().GetBool("server.DumpRouterMap")
	Config.Server.ErrorLogEnabled = g.Cfg().GetBool("server.ErrorLogEnabled")
	Config.Server.AccessLogEnabled = g.Cfg().GetBool("server.AccessLogEnabled")
}

func init() {
	Config.Captcha.KeyLong = g.Cfg("captcha").GetInt("captcha.KeyLong")
	Config.Captcha.ImgWidth = g.Cfg("captcha").GetInt("captcha.ImgWidth")
	Config.Captcha.ImgHeight = g.Cfg("captcha").GetInt("captcha.ImgHeight")
}

func init() {
	Config.Database.Host = g.Cfg().GetString("database.default.Host")
	Config.Database.Port = g.Cfg().GetString("database.default.Port")
	Config.Database.User = g.Cfg().GetString("database.default.User")
	Config.Database.Pass = g.Cfg().GetString("database.default.Pass")
	Config.Database.Name = g.Cfg().GetString("database.default.Name")
	Config.Database.Type = g.Cfg().GetString("database.default.Type")
	Config.Database.Role = g.Cfg().GetString("database.default.Role")
	Config.Database.Debug = g.Cfg().GetBool("database.default.Debug")
	Config.Database.Prefix = g.Cfg().GetString("database.default.Prefix")
	Config.Database.DryRun = g.Cfg().GetBool("database.default.DryRun")
	Config.Database.Weight = g.Cfg().GetInt("database.default.Weight")
	Config.Database.Charset = g.Cfg().GetString("database.default.Charset")
	Config.Database.LinkInfo = g.Cfg().GetString("database.default.LinkInfo")
	Config.Database.MaxIdleConnCount = g.Cfg().GetInt("database.default.MaxIdleConnCount")
	Config.Database.MaxOpenConnCount = g.Cfg().GetInt("database.default.MaxOpenConnCount")
	Config.Database.MaxConnLifetime = g.Cfg().GetDuration("database.default.MaxConnLifetime")
}

func init() {
	Config.DatabaseLogger.Path = g.Cfg().GetString("database.logger.Path")
	Config.DatabaseLogger.Level = g.Cfg().GetString("database.logger.Level")
	Config.DatabaseLogger.Stdout = g.Cfg().GetBool("database.logger.Stdout")
}
