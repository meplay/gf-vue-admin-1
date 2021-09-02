package config

import config "github.com/flipped-aurora/gf-vue-admin/library/config/gorm"

type Config struct {
	Zap   Zap         `mapstructure:"zap" json:"zap" yaml:"zap"`       // zap
	Gorm  config.Gorm `mapstructure:"gorm" json:"gorm" yaml:"gorm"`    // gorm
	Redis Redis       `mapstructure:"redis" json:"redis" yaml:"redis"` // redis
}
