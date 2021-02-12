package boot

func Initialize(path ...string) {
	Viper.Initialize(path...)
	Gorm.Initialize()
	Redis.Initialize()
}
