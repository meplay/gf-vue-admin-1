package config

import (
	"fmt"
	"time"
)

type Minio struct {
	Id       string `mapstructure:"id" json:"id" yaml:"id"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Token    string `mapstructure:"token" json:"token" yaml:"token"`
	Bucket   string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Endpoint string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`

	UseSsl bool `mapstructure:"use-ssl" json:"useSsl" yaml:"use-ssl"`
}
// Filename 文件名格式 自己可以改 建议保证唯一性
// Author [SliverHorn](https://github.com/SliverHorn)
func (m *Minio) Filename(filename string) string {
	folder := time.Now().Format("2006-01-02")
	return fmt.Sprintf("%s/%d%s", folder, time.Now().Unix(), filename)
}

func (m *Minio) Filepath(filename string) string {
	return m.Path + "/" + m.Bucket + "/" + filename
}