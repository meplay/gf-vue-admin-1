package global

import (
	"gf-vue-admin/library/config"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/os/gtime"
	"github.com/spf13/viper"
)

var (
	Viper  *viper.Viper
	Redis  *redis.Client
	Config config.Config
)

type Model struct {
	Id       uint        `orm:"id,primary"   json:"ID"`        // 自增ID
	CreateAt *gtime.Time `orm:"create_at"    json:"CreatedAt"` // 创建时间
	UpdateAt *gtime.Time `orm:"update_at"    json:"UpdatedAt"` // 更新时间
	DeleteAt *gtime.Time `orm:"delete_at"    json:"DeletedAt"` // 删除时间
}