package global

import (
	"github.com/flipped-aurora/gf-vue-admin/library/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"
)

var (
	Db     *gorm.DB
	Redis  *redis.Client
	Config config.Config
)

type Model struct {
	ID        uint           `gorm:"PrimaryKey"`     // 主键ID
	CreatedAt time.Time      `json:"CreatedAt"`      // 创建时间
	UpdatedAt time.Time      `json:"UpdatedAt"`      // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // 删除时间
}
