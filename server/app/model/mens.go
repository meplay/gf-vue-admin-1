package model

import (
	"server/app/model/parameters"

	"github.com/gogf/gf/os/gtime"
)

type AuthorityMenu struct {
	BaseMenu
	MenuId      string           `orm:"menu_id" json:"menuId"` //菜单ID
	AuthorityId string           `orm:"authority_id" json:"-"` // 角色ID
	Children    []*AuthorityMenu `orm:"children" json:"children"`
}

type BaseMenu struct {
	Id         uint                    `p:"ID" orm:"id,primary"   json:"ID"`               // 自增ID
	CreateAt   *gtime.Time             `p:"CreateAt" orm:"create_at"   json:"CreateAt"`    // 创建时间
	UpdateAt   *gtime.Time             `p:"UpdateAt" orm:"update_at"   json:"UpdateAt"`    // 更新时间
	DeleteAt   *gtime.Time             `p:"DeleteAt" orm:"delete_at"   json:"DeleteAt"`    // 删除时间
	MenuLevel  uint                    `p:"-" orm:"menu_level"   json:"-"`                 // 菜单等级(预留字段)
	ParentId   string                  `p:"parentId" orm:"parent_id"    json:"parentId"`   // 父菜单ID
	Path       string                  `p:"path" orm:"path"         json:"path"`           // 路由path
	Name       string                  `p:"name" orm:"name"         json:"name"`           // 路由name
	Hidden     bool                    `p:"hidden" orm:"hidden"       json:"hidden"`       // 是否在列表隐藏
	Component  string                  `p:"component" orm:"component"    json:"component"` // 前端文件路径
	Sort       int                     `p:"sort" orm:"sort"         json:"sort"`           // 排序标记
	Meta       `json:"meta"`           // 附加属性
	Children   []*BaseMenu             `orm:"children" json:"children"`
	Parameters []parameters.Parameters `json:"parameters"`
}

type Meta struct {
	Title       string `json:"title" orm:"title"        json:"title"`             // 菜单名
	Icon        string `json:"icon" orm:"icon"         json:"icon"`               // 菜单图标
	KeepAlive   int    `json:"keepAlive" orm:"keep_alive"   json:"keepAlive"`     // 是否缓存
	DefaultMenu bool   `json:"defaultMenu" orm:"default_menu" json:"defaultMenu"` // 是否是基础路由(开发中)
}
