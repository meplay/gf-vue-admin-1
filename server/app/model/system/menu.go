package model

import "gf-vue-admin/library/global"

type BaseMenu struct {
	global.Model
	Path      string `json:"path" gorm:"comment:路由path"`
	Name      string `json:"name" gorm:"comment:路由name"`
	ParentId  string `json:"parentId" gorm:"comment:父菜单ID"`
	Component string `json:"component" gorm:"comment:对应前端文件路径"`
	Sort      int    `json:"sort" gorm:"comment:排序标记"`
	MenuLevel uint   `json:"-"`
	Hidden    bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`

	Meta `json:"meta" gorm:"comment:附加属性"`

	Children    []BaseMenu          `json:"children" gorm:"-"`
	Parameters  []BaseMenuParameter `json:"parameters" gorm:"many2many:menus_parameters;foreignKey:ID;joinForeignKey:BaseMenuID;References:BaseMenuID;JoinReferences:ParameterID"`
	Authorities []Authority         `json:"authoritys" gorm:"many2many:authorities_menus;foreignKey:ID;joinForeignKey:BaseMenuID;References:AuthorityId;JoinReferences:AuthorityId"`
}

func (b *BaseMenu) TableName() string {
	return "menus"
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"`
	Icon        string `json:"icon" gorm:"comment:菜单图标"`
	Title       string `json:"title" gorm:"comment:菜单名"`
}

type BaseMenuParameter struct {
	global.Model
	Key        string `json:"key" gorm:"comment:地址栏携带参数的key"`
	Type       string `json:"type" gorm:"comment:地址栏携带参数为params还是query"`
	Value      string `json:"value" gorm:"comment:地址栏携带参数的值"`
	BaseMenuID uint   `json:"base_menu_id" gorm:"comment:base_menu的id"`
}

func (b *BaseMenuParameter) TableName() string {
	return "menu_parameter"
}
