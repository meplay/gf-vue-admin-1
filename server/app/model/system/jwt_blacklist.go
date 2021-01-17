package model

import "github.com/gogf/gf/os/gtime"

type JwtBlacklist struct {
	Id       uint        `orm:"id,primary" json:"id"`        // 自增ID
	CreateAt *gtime.Time `orm:"create_at"  json:"create_at"` // 更新时间
	UpdateAt *gtime.Time `orm:"update_at"  json:"update_at"` // 更新时间
	DeleteAt *gtime.Time `orm:"delete_at"  json:"delete_at"` // 删除时间
	Jwt      string      `orm:"jwt"        json:"jwt"`       // jwt
}
