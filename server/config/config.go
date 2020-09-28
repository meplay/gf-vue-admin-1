package config

type Server struct {
	Qiniu  Qiniu
	Minio  Minio
	AliYun AliYun
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
