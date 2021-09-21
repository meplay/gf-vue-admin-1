//go:build mysql
// +build mysql

package boot

import (
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var _ interfaces.Gorm = (*_mysql)(nil)

var DbResolver = new(_mysql)

type _mysql struct {
	dsn string
}

// GetSources 获取主库的 gorm.Dialector 切片对象
// Author [SliverHorn](https://github.com/SliverHorn)
func (m *_mysql) GetSources() (directories []gorm.Dialector) {
	length := len(global.Config.Gorm.Dsn.Sources)
	directories = make([]gorm.Dialector, 0, length)
	for i := 0; i < length; i++ {
		if !global.Config.Gorm.Dsn.Sources[i].IsEmpty() {
			dsn := global.Config.Gorm.Dsn.Sources[i].GetDsn(global.Config.Gorm.Config)
			if i == 0 {
				m.dsn = dsn
			}
			directories = append(directories, mysql.Open(dsn))
		} else {
			continue
		}
	}
	return directories
}

// GetReplicas 获取从库库的 gorm.Dialector 切片对象
// Author [SliverHorn](https://github.com/SliverHorn)
func (m *_mysql) GetReplicas() (directories []gorm.Dialector) {
	length := len(global.Config.Gorm.Dsn.Replicas)
	directories = make([]gorm.Dialector, 0, length)
	for i := 0; i < length; i++ {
		if !global.Config.Gorm.Dsn.Replicas[i].IsEmpty() {
			dsn := global.Config.Gorm.Dsn.Replicas[i].GetDsn(global.Config.Gorm.Config)
			directories = append(directories, mysql.Open(dsn))
		} else {
			continue
		}
	}
	return directories
}

// GetResolver 通过主库与从库的链接组装 gorm.Plugin
// Author [SliverHorn](https://github.com/SliverHorn)
func (m *_mysql) GetResolver() gorm.Plugin {
	sources := m.GetSources()
	resolver := dbresolver.Register(dbresolver.Config{
		Sources:  sources,
		Replicas: m.GetReplicas(),
		Policy:   dbresolver.RandomPolicy{}, // sources/replicas 负载均衡策略
	})
	resolver.SetMaxIdleConns(global.Config.Gorm.GetMaxOpenConnes())
	resolver.SetMaxOpenConns(global.Config.Gorm.GetMaxOpenConnes())
	resolver.SetConnMaxIdleTime(global.Config.Gorm.GetConnMaxIdleTime())
	resolver.SetConnMaxLifetime(global.Config.Gorm.GetConnMaxLifetime())
	return resolver
}

// GetGormDialector 获取数据库的 gorm.Dialector
// Author [SliverHorn](https://github.com/SliverHorn)
func (m *_mysql) GetGormDialector() gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN:                       m.dsn, // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: true,  // 根据版本自动配置
	})
}

func (m *_mysql) GetConfigPath() string {
	return "config/config.mysql.yaml"
}
