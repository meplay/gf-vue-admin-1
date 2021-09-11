package config

import config "github.com/flipped-aurora/gf-vue-admin/library/config/gorm"

type Config struct {
	Jwt     Jwt     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`             // jwt
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`             // zap
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`       // redis
	System  System  `mapstructure:"system" json:"system" yaml:"system"`    // system
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`    // casbin
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"` // captcha

	Gorm config.Gorm `mapstructure:"gorm" json:"gorm" yaml:"gorm"` // gorm
}
