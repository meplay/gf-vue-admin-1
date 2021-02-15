package request

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"gorm.io/gorm"
)

type UpdateBaseMenu struct {
	global.Model
	Path       string `json:"path" gorm:"comment:路由path"`
	Name       string `json:"name" gorm:"comment:路由name"`
	ParentId   string `json:"parentId" gorm:"comment:父菜单ID"`
	Component  string `json:"component" gorm:"comment:对应前端文件路径"`
	Sort       int    `json:"sort" gorm:"comment:排序标记"`
	MenuLevel  uint   `json:"-"`
	Hidden     bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`
	model.Meta `json:"meta" gorm:"comment:附加属性"`
	Parameters []model.MenuParameter `json:"parameters" gorm:"many2many:menus_parameters;foreignKey:ID;joinForeignKey:BaseMenuID;References:BaseMenuID;JoinReferences:ParameterID"`
}

func (u *UpdateBaseMenu) Update() func(db *gorm.DB) *gorm.DB {
	var update = make(map[string]interface{}, 10)
	update["keep_alive"] = u.KeepAlive
	update["default_menu"] = u.DefaultMenu
	update["parent_id"] = u.ParentId
	update["path"] = u.Path
	update["name"] = u.Name
	update["hidden"] = u.Hidden
	update["component"] = u.Component
	update["title"] = u.Title
	update["icon"] = u.Icon
	update["sort"] = u.Sort
	return func(db *gorm.DB) *gorm.DB {
		return db.Updates(update)
	}

}

