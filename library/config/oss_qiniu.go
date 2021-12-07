package config

import "github.com/qiniu/api.v7/v7/storage"

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`                                // 存储区域
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                          // 空间名称
	ImgPath       string `mapstructure:"img-path" json:"imgPath" yaml:"img-path"`                     // CDN加速域名
	UseHTTPS      bool   `mapstructure:"use-https" json:"useHttps" yaml:"use-https"`                  // 是否使用https
	AccessKey     string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`               // 秘钥AK
	SecretKey     string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`               // 秘钥SK
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"useCdnDomains" yaml:"use-cdn-domains"` // 上传是否使用CDN上传加速
}

func (q *Qiniu) GetConfig() *storage.Config {
	cfg := storage.Config{
		Zone:          q.GetZone(),
		UseHTTPS:      q.UseHTTPS,
		UseCdnDomains: q.UseCdnDomains,
	}
	return &cfg
}

func (q *Qiniu) GetZone() *storage.Region {
	switch q.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuaDong":
		return &storage.ZoneHuadong
	case "ZoneHuaBei":
		return &storage.ZoneHuabei
	case "ZoneHuaNan":
		return &storage.ZoneHuanan
	case "ZoneBeiMei":
		return &storage.ZoneBeimei
	case "ZoneXinJiaPo":
		return &storage.ZoneXinjiapo
	}
	return &storage.ZoneHuadong
}
