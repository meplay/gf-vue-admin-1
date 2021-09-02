package config

type Config struct {
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
}
