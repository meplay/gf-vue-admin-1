package config

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
}

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	ImgPath       string `mapstructure:"img-path" json:"imgPath" yaml:"img-path"`
	UseHTTPS      bool   `mapstructure:"use-https" json:"useHttps" yaml:"use-https"`
	AccessKey     string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`
	SecretKey     string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"useCdnDomains" yaml:"use-cdn-domains"`
}

type Minio struct {
	Id       string `mapstructure:"id" json:"id" yaml:"id"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Token    string `mapstructure:"token" json:"token" yaml:"token"`
	Bucket   string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	UseSsl   bool   `mapstructure:"use-ssl" json:"useSsl" yaml:"use-ssl"`
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Endpoint string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
}

type Aliyun struct {
	Path             string `mapstructure:"path" json:"path" yaml:"path"`
	Bucket           string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	AclType          string `mapstructure:"acl-type" json:"aclType" yaml:"acl-type"`
	Endpoint         string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyID      string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`
	SecretAccessKey  string `mapstructure:"secret-access-key" json:"secretAccessKey" yaml:"secret-access-key"`
	StorageClassType string `mapstructure:"storage-class-type" json:"storageClassType" yaml:"storage-class-type"`
}
