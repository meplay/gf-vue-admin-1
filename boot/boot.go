package boot

import (
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	boot "github.com/flipped-aurora/gf-vue-admin/boot/gorm"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
)

func Initialize() {
	Viper.Initialize()
	Zap.Initialize()
	Gorm.Initialize(boot.DbResolver)
	if global.Db != nil {
		system.JwtBlacklist.LoadJwt()
	}
	if global.Config.System.UseMultipoint {
		Redis.Initialize()
	}
}
