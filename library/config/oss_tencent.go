package config

type Tencent struct {
	Bucket     string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Region     string `mapstructure:"region" json:"region" yaml:"region"`
	BaseURL    string `mapstructure:"base-url" json:"baseURL" yaml:"base-url"`
	SecretID   string `mapstructure:"secret-id" json:"secretID" yaml:"secret-id"`
	SecretKey  string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	PathPrefix string `mapstructure:"path-prefix" json:"pathPrefix" yaml:"path-prefix"`
}

func (t *Tencent) Filename(filename string) string {
	return t.PathPrefix + "/" + filename
}

func (t *Tencent) Filepath(filename string) string {
	return t.BaseURL + "/" + filename
}
