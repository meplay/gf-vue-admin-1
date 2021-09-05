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
	ID        uint           `gorm:"PrimaryKey;column:id;comment:主键ID"`
	CreatedAt time.Time      `json:"CreatedAt" gorm:"column:created_at;comment:创建时间"`
	UpdatedAt time.Time      `json:"UpdatedAt" gorm:"column:updated_at;comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;column:deleted_at;comment:删除时间"`
}
