package boot

import (
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type dbResolver struct{}

// GetSources 获取主库的 gorm.Dialector 切片对象
// Author [SliverHorn](https://github.com/SliverHorn)
func (d *dbResolver) GetSources() (directories []gorm.Dialector) {
	length := len(global.Config.Gorm.Dsn.Sources)
	directories = make([]gorm.Dialector, 0, length)
	for i := 0; i < length; i++ {
		if !global.Config.Gorm.Dsn.Sources[i].IsEmpty() {
			dsn := global.Config.Gorm.Dsn.Sources[i].GetDsn(global.Config.Gorm.Config)
			directories = append(directories, DbResolver.GetGormDialector(dsn))
		} else {
			continue
		}
	}
	return directories
}

// GetReplicas 获取从库库的 gorm.Dialector 切片对象
// Author [SliverHorn](https://github.com/SliverHorn)
func (d *dbResolver) GetReplicas() (directories []gorm.Dialector) {
	length := len(global.Config.Gorm.Dsn.Replicas)
	directories = make([]gorm.Dialector, 0, length)
	for i := 0; i < length; i++ {
		if !global.Config.Gorm.Dsn.Replicas[i].IsEmpty() {
			dsn := global.Config.Gorm.Dsn.Replicas[i].GetDsn(global.Config.Gorm.Config)
			directories = append(directories, DbResolver.GetGormDialector(dsn))
		} else {
			continue
		}
	}
	return directories
}

// GetResolver 通过主库与从库的链接组装 gorm.Plugin
// Author [SliverHorn](https://github.com/SliverHorn)
func (d *dbResolver) GetResolver() gorm.Plugin {
	sources := d.GetSources()
	resolver := dbresolver.Register(dbresolver.Config{
		Sources:  sources,
		Replicas: d.GetReplicas(),
		Policy:   dbresolver.RandomPolicy{}, // sources/replicas 负载均衡策略
	})
	resolver.SetMaxIdleConns(global.Config.Gorm.GetMaxOpenConnes())
	resolver.SetMaxOpenConns(global.Config.Gorm.GetMaxOpenConnes())
	resolver.SetConnMaxIdleTime(global.Config.Gorm.GetConnMaxIdleTime())
	resolver.SetConnMaxLifetime(global.Config.Gorm.GetConnMaxLifetime())
	return resolver
}
