package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	ErrorToEmail  bool   `mapstructure:"error-to-email" json:"errorToEmail" yaml:"error-to-email"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"UseMultipoint" yaml:"use-multipoint"`
}