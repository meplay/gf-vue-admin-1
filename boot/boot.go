package boot

import boot "github.com/flipped-aurora/gf-vue-admin/boot/gorm"

func Initialize() {
	Viper.Initialize()
	Zap.Initialize()
	Gorm.Initialize(boot.DbResolver)
	Redis.Initialize()
	Routers.Initialize()
}
