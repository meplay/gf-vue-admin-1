package model

import (
	"time"
)

type ServerConfig struct {
	Jwt            Jwt            `json:"jwt"`
	Oss            Oss            `json:"oss"`
	Redis          Redis          `json:"redis"`
	Email          Email          `json:"email"`
	Casbin         Casbin         `json:"casbin"`
	Logger         Logger         `json:"logger"`
	Server         Server         `json:"server"`
	System         System         `json:"system"`
	Captcha        Captcha        `json:"captcha"`
	Database       Database       `json:"database"`
	DatabaseLogger DatabaseLogger `json:"database_logger"`
}

type Jwt struct {
	ExpiresAt  int    `json:"expires_at"`
	RefreshAt  int    `json:"refresh_at"`
	SigningKey string `json:"signing_key"`
}

type Oss struct {
	Local  Local  `json:"local"`
	Qiniu  Qiniu  `json:"qiniu"`
	Minio  Minio  `json:"minio"`
	Aliyun Aliyun `json:"aliyun"`
}

type Redis struct {
	Default string `json:"default"`
	Cache   string `json:"cache"`
}

type Email struct {
	To       string `json:"to"`
	Port     int    `json:"port"`
	From     string `json:"from"`
	Host     string `json:"host"`
	IsSsl    bool   `json:"is_ssl"`
	Secret   string `json:"secret"`
	Nickname string `json:"nickname"`
}

type Casbin struct {
	ModelPath string `json:"model_path"`
}

type Logger struct {
	Path   string `json:"path"`
	Level  string `json:"level"`
	Stdout bool   `json:"stdout"`
}

type Server struct {
	LogPath          string `json:"log_path"`
	Address          string `json:"address"`
	DumpRouterMap    bool   `json:"dump_router_map"`
	ErrorLogEnabled  bool   `json:"error_log_enabled"`
	AccessLogEnabled bool   `json:"access_log_enabled"`
}

type System struct {
	Db            string `json:"db"`
	Env           string `json:"env"`
	OssType       string `json:"oss_type"`
	ErrorToEmail  bool   `json:"error_to_email"`
	UseMultipoint bool   `json:"use_multipoint"`
}

type Captcha struct {
	KeyLong   int `json:"key_long"`
	ImgWidth  int `json:"img_width"`
	ImgHeight int `json:"img_height"`
}

type Local struct {
	LocalPath string `json:"local_path"`
}

type Qiniu struct {
	Zone          string `json:"zone"`
	Bucket        string `json:"bucket"`
	ImgPath       string `json:"img_path"`
	UseHTTPS      bool   `json:"use_https"`
	AccessKey     string `json:"access_key"`
	SecretKey     string `json:"secret_key"`
	UseCdnDomains bool   `json:"use_cdn_domains"`
}

type Minio struct {
	Id       string `json:"id"`
	Path     string `json:"path"`
	Token    string `json:"token"`
	Bucket   string `json:"bucket"`
	UseSsl   bool   `json:"use_ssl"`
	Secret   string `json:"secret"`
	Endpoint string `json:"endpoint"`
}

type Aliyun struct {
	Path             string `json:"path"`
	Bucket           string `json:"bucket"`
	ACLType          string `json:"acl_type"`
	Endpoint         string `json:"endpoint"`
	AccessKeyID      string `json:"access_key_id"`
	SecretAccessKey  string `json:"secret_access_key"`
	StorageClassType string `json:"storage_class_type"`
}

type Database struct {
	Host             string        `json:"host"`
	Port             string        `json:"port"`
	User             string        `json:"user"`
	Pass             string        `json:"pass"`
	Name             string        `json:"name"`
	Type             string        `json:"type"`
	Role             string        `json:"role"`
	Debug            bool          `json:"debug"`
	Prefix           string        `json:"prefix"`
	DryRun           bool          `json:"dry_run"`
	Weight           int           `json:"weight"`
	Charset          string        `json:"charset"`
	LinkInfo         string        `json:"link_info"`
	MaxIdleConnCount int           `json:"max_idle_conn_count"`
	MaxOpenConnCount int           `json:"max_open_conn_count"`
	MaxConnLifetime  time.Duration `json:"max_conn_lifetime"`
}

type DatabaseLogger struct {
	Path   string `json:"path"`
	Level  string `json:"level"`
	Stdout bool   `json:"stdout"`
}
