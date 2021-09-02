package boot

func Initialize() {
	Viper.Initialize()
	Zap.Initialize()
	Gorm.Initialize()
	Redis.Initialize()
}
