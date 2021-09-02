package global

import (
	"github.com/flipped-aurora/gf-vue-admin/library/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Db     *gorm.DB
	Redis  *redis.Client
	Config config.Config
)
