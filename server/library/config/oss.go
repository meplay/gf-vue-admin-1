package config

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
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
