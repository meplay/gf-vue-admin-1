package service

import (
	"errors"
	model "flipped-aurora/gf-vue-admin/server/app/model/system"
	"flipped-aurora/gf-vue-admin/server/app/model/system/request"
	"flipped-aurora/gf-vue-admin/server/app/service/system/internal"
	"flipped-aurora/gf-vue-admin/server/library/global"
	"flipped-aurora/gf-vue-admin/server/library/response"
	"gorm.io/gorm"
)

var Menu = new(menu)

type menu struct {
	_menu             model.Menu
	_parameter        model.MenuParameter
	_menusParameters  model.MenusParameters
	_authoritiesMenus model.AuthoritiesMenus
}

// Create 添加基础路由
// Author [Aizen1172](https://github.com/Aizen1172)
func (m *menu) Create(menu *model.Menu) error {
	return global.Db.Create(&menu).Error
}

// First 返回当前选中menu
// Author [Aizen1172](https://github.com/Aizen1172)
func (m *menu) First(info *request.GetById) (menu *model.Menu, err error) {
	var entity model.Menu
	err = global.Db.Where("id = ?", info.Id).First(&entity).Error
	var params = internal.Menu.GetMenusParameters(entity.ID)
	if params != nil {
		entity.Parameters = *params
	}
	return &entity, err
}

// Delete 删除基础路由
// Author [Aizen1172](https://github.com/Aizen1172)
func (m *menu) Delete(info *request.GetById) error {
	var entity model.Menu
	if errors.Is(global.Db.Where("parent_id", info.Id).First(&entity).Error, gorm.ErrRecordNotFound) {
		return response.ErrorHasChildrenMenu
	}
	var authorities = internal.Menu.GetAuthoritiesMenus(entity.ID)
	if authorities != nil {
		entity.Authorities = *authorities
	}
	if err := global.Db.Delete(&model.MenuParameter{}, info.Id).Error; err != nil {
		return err
	}
	if err := global.Db.Where("menu_id", info.Id).Delete(&model.AuthoritiesMenus{}).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新路由
// Author [Aizen1172](https://github.com/Aizen1172)
func (m *menu) Update(info *request.UpdateMenu) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		var entity model.Menu
		if err := tx.Where("id = ?", info.ID).First(&entity).Error; err != nil {
			return err
		}
		if entity.Name != info.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", info.ID, info.Name).First(&model.Menu{}).Error, gorm.ErrRecordNotFound) {
				return response.ErrorUpdateMenuName
			}
		}
		if err := global.Db.Delete(&model.MenuParameter{}, info.ID).Error; err != nil {
			return err
		}
		if err := tx.Where("id", info.ID).Updates(info.Update()).Error; err != nil {
			return response.ErrorUpdateMenu
		}
		if err := tx.Unscoped().Delete(&model.MenusParameters{}, info.ID).Error; err != nil {
			return err
		}
		for _, parameter := range info.Parameters {
			_entity := &model.MenusParameters{MenuId: info.ID, ParameterId: parameter.ID}
			if err := tx.Create(&_entity).Error; err != nil {
				return response.ErrorCreateParameters
			}
		}
		return nil
	})
}

// GetList 获取路由分页
// Author [Aizen1172](https://github.com/Aizen1172)
func (m *menu) GetList() (list interface{}, total int, err error) {
	var menus []model.Menu
	var treeMap = internal.Menu.GetTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		internal.Menu.GetChildrenList(&menus[i], treeMap)
	}
	return menus, total, err
}

func (m *menu) GetTree() *[]model.Menu {
	tree := internal.Menu.GetTreeMap()
	menus := tree["0"]
	for i := 0; i < len(menus); i++ {
		internal.Menu.GetChildrenList(&menus[i], tree)
	}
	return &menus
}
