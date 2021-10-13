//go:build mysql
// +build mysql

package boot

import (
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
)

var _ interfaces.Gorm = (*_mysql)(nil)

var DbResolver = new(_mysql)

type _mysql struct {
	dbResolver
}

// GetGormDialector 获取数据库的 gorm.Dialector
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) GetGormDialector(dsn string) gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN:                       dsn,  // DSN data source name
		DefaultStringSize:         191,  // string 类型字段的默认长度
		SkipInitializeWithVersion: true, // 根据版本自动配置
	})
}

func (g *_gorm) GetConfigPath() string {
	return "config/config.mysql.yaml"
}
