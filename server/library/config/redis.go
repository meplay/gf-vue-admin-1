package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Address  string `mapstructure:"address" json:"address" yaml:"address"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
