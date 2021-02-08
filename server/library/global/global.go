package global

import (
	"gf-vue-admin/library/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	Config config.Config
	Redis  *redis.Client
	Viper  *viper.Viper
)
