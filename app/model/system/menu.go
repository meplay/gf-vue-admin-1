package system

import "github.com/flipped-aurora/gf-vue-admin/library/global"

type Menu struct {
	global.Model
	Sort        int    `json:"sort" gorm:"column:sort;comment:排序标记"`
	Hidden      bool   `json:"hidden" gorm:"column:hidden;comment:是否在列表隐藏"`
	MenuLevel   uint   `json:"-"`
	Path        string `json:"path" gorm:"column:path;comment:路由path"`
	Name        string `json:"name" gorm:"column:name;comment:路由name"`
	Component   string `json:"component" gorm:"column:component;comment:对应前端文件路径"`
	ParentId    string `json:"parentId" gorm:"column:parent_id;comment:父菜单ID"`
	Meta        `json:"meta" gorm:"comment:附加属性"`
	Children    []Menu          `json:"children" gorm:"-"`
	Parameters  []MenuParameter `json:"parameters" gorm:"many2many:system_menus_parameters;foreignKey:ID;joinForeignKey:MenuID;References:MenuID;JoinReferences:ParameterID"`
	Authorities []Authority     `json:"authoritys" gorm:"many2many:system_authorities_menus;foreignKey:ID;joinForeignKey:MenuID;References:AuthorityId;JoinReferences:AuthorityId"`
}

type Meta struct {
	Icon        string `json:"icon" gorm:"column:icon;comment:菜单图标"`
	Title       string `json:"title" gorm:"column:title;comment:菜单名"`
	CloseTab    bool   `json:"closeTab" gorm:"column:close_tab;comment:自动关闭tab"`
	KeepAlive   bool   `json:"keepAlive" gorm:"column:keep_alive;comment:是否缓存"`
	DefaultMenu bool   `json:"defaultMenu" gorm:"column:default_menu;comment:是否是基础路由(开发中)"`
}

func (m *Menu) TableName() string {
	return "system_menus"
}
