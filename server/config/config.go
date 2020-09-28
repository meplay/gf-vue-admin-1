package config

type Server struct {
	System System
	Local  Local
	Qiniu  Qiniu
	Minio  Minio
	AliYun AliYun
}

type System struct {
	Db            string
	Env           string
	OssType       string // local:本地,qiniu:七牛云存储,minio:minio存储,aliyun:阿里云存储
	UseMultipoint string
}

type Local struct {
	Path string
}

type Qiniu struct {
	Zone          string
	Bucket        string
	ImgPath       string
	UseHTTPS      bool
	AccessKey     string
	SecretKey     string
	UseCdnDomains bool
}

type Minio struct {
	Id       string
	Path     string
	Token    string
	Bucket   string
	Secret   string
	Endpoint string
	UseSSL   bool
}

type AliYun struct {
	Path             string
	Bucket           string
	Endpoint         string
	AccessKeyID      string
	SecretAccessKey  string
	StorageClassType string
}
