package global

import (
	"time"

	"github.com/flipped-aurora/gf-vue-admin/library/config"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	Db                 *gorm.DB
	Redis              *redis.Client
	Viper              *viper.Viper
	Config             config.Config
	JwtCache           local_cache.Cache
	ConcurrencyControl = &singleflight.Group{}
)

type Model struct {
	ID        uint           `gorm:"PrimaryKey;column:id;comment:主键ID" example:"7"`
	CreatedAt time.Time      `json:"CreatedAt" gorm:"column:created_at;comment:创建时间" example:"创建时间"`
	UpdatedAt time.Time      `json:"UpdatedAt" gorm:"column:updated_at;comment:更新时间" example:"更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;column:deleted_at;comment:删除时间" example:"删除时间"`
}
