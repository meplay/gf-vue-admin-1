package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	OrmType       string `mapstructure:"orm-type" json:"ormType" yaml:"orm-type"`
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	ErrorToEmail  bool   `mapstructure:"error-to-email" json:"errorToEmail" yaml:"error-to-email"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"UseMultipoint" yaml:"use-multipoint"`
}
