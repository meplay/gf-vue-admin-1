package service

import (
	"gf-vue-admin/app/api/request"
	"gf-vue-admin/app/api/response"
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/app/service/system/internal"
	"github.com/gogf/gf/frame/g"
)

var Menu = new(menu)

type menu struct {
	_menu      model.Menu
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 返回当前选中menu
func (m *menu) First(info *request.GetById) (menu *model.Menu, err error) {
	var entity model.Menu
	err = g.DB().Table(m._menu.TableName()).Where(info.Condition()).Struct(&entity)
	entity.Parameters = *internal.Menu.GetMenusParameters(entity.ID)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除基础路由
func (m *menu) Delete(info *request.GetById) error {
	//var menu model.Menu
	//var err = global.Db.Preload("Parameters").Preload("Authorities").Where("parent_id = ?", info.ID).First(&menu).Error
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	if err = global.Db.Delete(&model.MenuParameter{}, "base_menu_id = ?", info.ID).Error; err == nil && len(menu.Authorities) > 0 {
	//		err = global.Db.Model(&menu).Association("Authoritys").Delete(&menu.Authorities)
	//	}
	//	return err
	//}
	return response.ErrorHasChildrenMenu
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 更新路由
func (m *menu) Update(info *request.UpdateBaseMenu) error {
	//var oldMenu model.Menu
	//return global.Db.Transaction(func(tx *gorm.DB) error {
	//	db := tx.Where("id = ?", info.ID).Find(&oldMenu)
	//	if oldMenu.Name != info.Name {
	//		if !errors.Is(tx.Where("id <> ? AND name = ?", info.ID, info.Name).First(&model.Menu{}).Error, gorm.ErrRecordNotFound) {
	//			return model.ErrorUpdateBaseMenuName
	//		}
	//	}
	//	if err := tx.Unscoped().Delete(&model.MenuParameter{}, "base_menu_id = ?", info.ID).Error; err != nil {
	//		zap.L().Debug(err.Error())
	//		return err
	//	}
	//	if len(info.Parameters) > 0 {
	//		for i, _ := range info.Parameters {
	//			info.Parameters[i].BaseMenuID = info.ID
	//		}
	//		if err := tx.Create(&info.Parameters).Error; err != nil {
	//			return model.ErrorCreateParameters
	//		}
	//	}
	//
	//	if err := db.Scopes(info.Update()).Error; err != nil {
	//		return model.ErrorUpdate
	//	}
	//	return nil
	//})
	return nil
}
