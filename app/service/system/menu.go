package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
)

var Menu = new(_menu)

type _menu struct{}

// Create 添加菜单
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) Create(info *request.MenuCreate) error {
	if !errors.Is(global.Db.Where("name = ?", info.Name).First(&system.Menu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name!")
	}
	return global.Db.Create(&info.Menu).Error
}

// Delete 删除基础路由
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) Delete(info *common.GetByID) (err error) {
	err = global.Db.Preload("Parameters").Where("parent_id = ?", info.ID).First(&system.Menu{}).Error
	if err != nil {
		var menu system.Menu
		db := global.Db.Preload("SysAuthoritys").Where("id = ?", info.ID).First(&menu).Delete(&menu)
		err = global.Db.Delete(&system.MenuParameter{}, "menu_id = ?", info.ID).Error
		if len(menu.Authorities) > 0 {
			if err = global.Db.Model(&menu).Association("Authoritys").Delete(&menu.Authorities); err != nil {
				return errors.Wrap(err, "删除菜单绑定的角色失败!")
			}
		} else {
			err = db.Error
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

// First 返回当前选中menu
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) First(info *common.GetByID) (*system.Menu, error) {
	var entity system.Menu
	err := global.Db.Preload("Parameters").Where("id = ?", info.ID).First(&entity).Error
	return &entity, err
}

// Update 更新路由
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) Update(info *request.MenuUpdate) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		var entity system.Menu
		if err := tx.Where("id = ?", info.ID).First(&entity).Error; err != nil {
			return errors.Wrap(err, "菜单信息获取失败!")
		}
		if entity.Name != info.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", info.ID, info.Name).First(&system.Menu{}).Error, gorm.ErrRecordNotFound) {
				return errors.New("存在相同name修改失败!")
			}
		}
		if err := tx.Unscoped().Delete(&system.MenuParameter{}, "menu_id = ?", info.ID).Error; err != nil {
			return errors.Wrap(err, "更新菜单参数信息失败!")
		}
		if len(info.Parameters) > 0 {
			for k := range info.Parameters {
				info.Parameters[k].MenuID = info.ID
			}
			if err := tx.Create(&info.Parameters).Error; err != nil {
				return errors.Wrap(err, "新增菜单参数信息失败!")
			}
		}
		update := info.Update()
		if err := tx.Updates(update).Error; err != nil {
			return errors.Wrap(err, "更新菜单信息失败!")
		}
		return nil
	})
}

// GetList 获取全部路由
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) GetList() (list []system.Menu, total int64, err error) {
	treeMap, _err := s.getTreeMap()
	if _err != nil {
		return nil, 0, _err
	}
	list = treeMap["0"]
	for i := 0; i < len(list); i++ {
		if err = s.getChildrenList(&list[i], treeMap); err != nil {
			return list, total, errors.Wrap(err, "查找子菜单数据失败!")
		}
	}
	return list, total, err
}

// GetTree 获取动态菜单树
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) GetTree() (list []system.Menu, err error) {
	treeMap, _err := s.getTreeMap()
	if _err != nil {
		return nil, _err
	}
	list = treeMap["0"]
	for i := 0; i < len(list); i++ {
		if err = s.getChildrenList(&list[i], treeMap); err != nil {
			return list, errors.Wrap(err, "查找子菜单数据失败!")
		}
	}
	return list, err
}

// AddMenuAuthority 为角色增加menu树
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) AddMenuAuthority(info *request.AddMenuAuthority) error {
	entity := info.ToAuthoritySetMenu()
	return Authority.SetAuthorityMenu(&entity)
}

// getTreeMap 获取路由总树map
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) getTreeMap() (treeMap map[string][]system.Menu, err error) {
	var entities []system.Menu
	if err = global.Db.Order("sort").Preload("Parameters").Find(&entities).Error; err != nil {
		return nil, errors.Wrap(err, "查找菜单数据失败!")
	}
	treeMap = make(map[string][]system.Menu, len(entities))
	for _, v := range entities {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, nil
}

// getChildrenList 获取菜单的子菜单
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) getChildrenList(menu *system.Menu, treeMap map[string][]system.Menu) error {
	if value, ok := treeMap[strconv.Itoa(int(menu.ID))]; ok {
		menu.Children = value
		for i := 0; i < len(menu.Children); i++ {
			if err := s.getChildrenList(&menu.Children[i], treeMap); err != nil {
				return err
			}
		}
		return nil
	} else {
		return nil
	}
}
