package boot

func Initialize() {
	Viper.Initialize()
	Zap.Initialize()
	//Gorm.Initialize(boot.DbResolver)
	//Redis.Initialize()
}
