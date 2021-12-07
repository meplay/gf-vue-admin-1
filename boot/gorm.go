package boot

import (
	"os"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gf-vue-admin/app/model/example"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	boot "github.com/flipped-aurora/gf-vue-admin/boot/gorm"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Gorm = new(_gorm)

type _gorm struct{}

func (g *_gorm) Initialize(i interfaces.Gorm) {
	resolver := boot.DbResolver.GetResolver()
	db, err := gorm.Open(i.GetGormDialector(global.Config.Gorm.Dsn.GetDefaultDsn(global.Config.Gorm.Config)), boot.Gorm.GenerateConfig())
	if err != nil {
		zap.L().Error("gorm链接数据库失败!", zap.Error(err))
		os.Exit(0)
	}
	err = db.Use(resolver)
	if err != nil {
		zap.L().Error("gorm链接集群失败!", zap.Error(err))
		os.Exit(0)
	}
	global.Db = db
	if global.Config.Gorm.AutoMigrate {
		go func() {
			err = global.Db.AutoMigrate(
				new(system.Api),
				new(system.User),
				new(system.Menu),
				new(system.Authority),
				new(system.Dictionary),
				new(system.MenuParameter),
				new(system.AutoCodeHistory),
				new(system.OperationRecord),
				new(system.DictionaryDetail),

				new(gormadapter.CasbinRule),

				new(example.Customer),
				// Code generated by gf-vue-admin Begin; DO NOT EDIT.
				// Code generated by gf-vue-admin End; DO NOT EDIT.
			)
			if err != nil {
				zap.L().Error("gorm注册表失败!", zap.Error(err))
				os.Exit(0)
			}
			zap.L().Info("gorm注册表成功!")
		}()
	}
	sql, _err := db.DB()
	if _err != nil {
		zap.L().Error("gorm设置数据库最大连接数失败!", zap.Error(_err))
		return
	}
	sql.SetMaxIdleConns(global.Config.Gorm.GetMaxIdleConnes())
	sql.SetMaxOpenConns(global.Config.Gorm.GetMaxOpenConnes())
}
