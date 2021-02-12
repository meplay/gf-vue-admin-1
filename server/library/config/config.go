package config

type Config struct {
	Jwt     Jwt     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Oss     Oss     `mapstructure:"oss" json:"oss" yaml:"oss"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`

	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
