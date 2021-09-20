package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

// GetAuthorityMenu 查看当前角色树
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authorityMenu) GetAuthorityMenu(info *common.GetAuthorityId) (menus []system.AuthorityMenu, err error) {
	err = global.Db.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	return menus, err
}

// GetMenuTree 获取动态菜单树
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authorityMenu) GetMenuTree(authorityId string) ([]system.AuthorityMenu, error) {
	treeMap, err := s.getMenuTreeMap(authorityId)
	if err != nil {
		return nil, err
	}
	menus := treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = s.getChildrenList(&menus[i], treeMap)
	}
	return menus, err
}

// getMenuTreeMap 获取路由总树map
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authorityMenu) getMenuTreeMap(authorityId string) (treeMap map[string][]system.AuthorityMenu, err error) {
	var entities []system.AuthorityMenu
	if err = global.Db.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&entities).Error; err != nil {
		return nil, errors.Wrap(err, "视图数据获取失败!")
	}
	treeMap = make(map[string][]system.AuthorityMenu, len(entities))
	for _, v := range entities {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, nil
}

// getChildrenList 获取子菜单列表
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authorityMenu) getChildrenList(menu *system.AuthorityMenu, treeMap map[string][]system.AuthorityMenu) error {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		if err := s.getChildrenList(&menu.Children[i], treeMap); err != nil {
			return err
		}
	}
	return nil
}
