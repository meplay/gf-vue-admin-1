package config

type Config struct {
	Jwt      Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email    Email    `mapstructure:"email" json:"email" yaml:"email"`
	Casbin   Casbin   `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	Captcha  Captcha  `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Uploader Uploader `mapstructure:"uploader" json:"uploader" yaml:"uploader"`

	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`

	// Oss
	Local  Local  `mapstructure:"local" json:"local" yaml:"local"`
	Qiniu  Qiniu  `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	Minio  Minio  `mapstructure:"minio" json:"minio" yaml:"minio"`
	Aliyun Aliyun `mapstructure:"aliyun" json:"aliyun" yaml:"aliyun"`
}
