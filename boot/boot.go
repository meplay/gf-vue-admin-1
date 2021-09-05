package boot

import "github.com/flipped-aurora/gf-vue-admin/app/router/middleware"

func Initialize() {
	Viper.Initialize()
	Zap.Initialize()
	middleware.Initialize()
	//Gorm.Initialize(boot.DbResolver)
	//Redis.Initialize()
}
