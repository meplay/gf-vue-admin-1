package config

import (
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/utils"
	"path"
	"strings"
	"time"
)

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件路径
}

// Filename 拼接新文件名
// Author [SliverHorn](https://github.com/SliverHorn)
func (l *Local) Filename(filename string) string {
	ext := path.Ext(filename)                 // 读取文件后缀
	name := strings.TrimSuffix(filename, ext) // 读取文件名并加密
	filename = utils.Encrypt.Md5([]byte(name))
	return filename + "_" + time.Now().Format("2006_01_02_15_04_05") + ext
}

// Filepath 拼接路径和文件名
// Author [SliverHorn](https://github.com/SliverHorn)
func (l *Local) Filepath(filename string) string {
	return global.Config.Local.Path + "/" + filename
}
