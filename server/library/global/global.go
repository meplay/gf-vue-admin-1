package global

import (
	"database/sql"
	"database/sql/driver"
	"flipped-aurora/gf-vue-admin/server/library/config"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/os/gtime"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

var (
	Db     *gorm.DB
	Viper  *viper.Viper
	Redis  *redis.Client
	Config config.Config
)

type _gTime gtime.Time

func (t *_gTime) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*t = _gTime(*gtime.NewFromTime(nullTime.Time))
	return
}

func (t _gTime) Value() (driver.Value, error) {
	y, m, d := gtime.Time(t).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, gtime.Time(t).Location()), nil
}

// GormDataType gorm common data type
func (t _gTime) GormDataType() string {
	return "date"
}


type Model struct {
	ID uint `orm:"id" json:"ID" gorm:"primaryKey;comment:主键ID"` // 自增ID

	CreatedAt time.Time `orm:"created_at" json:"CreatedAt" gorm:"column:created_at;comment:创建时间"` // 创建时间

	UpdatedAt time.Time `orm:"updated_at" json:"UpdatedAt" gorm:"column:updated_at;comment:更新时间"` // 更新时间

	DeletedAt gorm.DeletedAt `orm:"deleted_at" json:"-" gorm:"index;column:deleted_at;comment:删除时间"` // 删除时间
}
