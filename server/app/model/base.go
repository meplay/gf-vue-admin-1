package model

import "github.com/gogf/gf/os/gtime"

type BaseModel struct {
	ID       int         `json:"id" orm:"id"`               // 自增ID
	CreateAt *gtime.Time `json:"create_at" orm:"create_at"` // 创建时间
	UpdateAt *gtime.Time `json:"update_at" orm:"update_at"` // 更新时间
	DeleteAt *gtime.Time `json:"delete_at" orm:"delete_at"` // 删除时间
}
