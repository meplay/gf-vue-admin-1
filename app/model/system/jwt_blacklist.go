package system

import "github.com/flipped-aurora/gf-vue-admin/library/global"

type JwtBlacklist struct {
	global.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (j *JwtBlacklist) TableName() string {
	return "system_jwt_blacklist"
}
