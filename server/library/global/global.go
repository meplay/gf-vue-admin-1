package global

import (
	"gf-vue-admin/library/config"
	"github.com/go-redis/redis"
)

var (
	Config *config.Config
	Redis  *redis.Client
)
