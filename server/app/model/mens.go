package model

import (
	"github.com/gogf/gf/os/gtime"
)

type AuthorityMenu struct {
	BaseMenu
	MenuId      string           `orm:"menu_id" json:"menuId"` //菜单ID
	AuthorityId string           `orm:"authority_id" json:"-"` // 角色ID
	Children    []*AuthorityMenu `orm:"children" json:"children"`
}

type BaseMenu struct {
	Id        uint          `v:"ID" orm:"id,primary"   json:"ID"`               // 自增ID
	CreateAt  *gtime.Time   `v:"CreateAt" orm:"create_at"   json:"CreateAt"`    // 创建时间
	UpdateAt  *gtime.Time   `v:"UpdateAt" orm:"update_at"   json:"UpdateAt"`    // 更新时间
	DeleteAt  *gtime.Time   `v:"DeleteAt" orm:"delete_at"   json:"DeleteAt"`    // 删除时间
	MenuLevel uint          `v:"-" orm:"menu_level"   json:"-"`                 // 菜单等级(预留字段)
	ParentId  string        `v:"parentId" orm:"parent_id"    json:"parentId"`   // 父菜单ID
	Path      string        `v:"path" orm:"path"         json:"path"`           // 路由path
	Name      string        `v:"name" orm:"name"         json:"name"`           // 路由name
	Hidden    bool          `v:"hidden" orm:"hidden"       json:"hidden"`       // 是否在列表隐藏
	Component string        `v:"component" orm:"component"    json:"component"` // 前端文件路径
	Sort      int           `v:"sort" orm:"sort"         json:"sort"`           // 排序标记
	Meta      `json:"meta"` // 附加属性
	Children  []*BaseMenu   `orm:"children" json:"children"`
}

type Meta struct {
	Title       string `orm:"title"        json:"title"`       // 菜单名
	Icon        string `orm:"icon"         json:"icon"`        // 菜单图标
	KeepAlive   int    `orm:"keep_alive"   json:"keepAlive"`   // 是否缓存
	DefaultMenu bool   `orm:"default_menu" json:"defaultMenu"` // 是否是基础路由(开发中)
}
