package config

type Captcha struct {
	KeyLong         int  `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`
	ImageWidth      int  `mapstructure:"image-width" json:"imageWidth" yaml:"image-width"`
	ImageHeight     int  `mapstructure:"image-height" json:"imageHeight" yaml:"image-height"`
	CaptchaInRedis  bool  `mapstructure:"captcha-in-redis" json:"captcha_in_redis" yaml:"captcha-in-redis"`
}
