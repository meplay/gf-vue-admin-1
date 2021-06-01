package internal

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct {
	_authorityMenu model.AuthorityMenu
}

// GetTreeMap 获取路由总树map
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authorityMenu) GetTreeMap(id string) map[string][]model.AuthorityMenu {
	entities := make([]model.AuthorityMenu, 0, 10)
	if err := global.Db.Where("authority_id", id).Order("sort").Find(&entities).Error;err != nil {
		g.Log().Error("获取 authority_menu 视图数据失败!", g.Map{"authority_id": id})
	}
	tree := make(map[string][]model.AuthorityMenu, len(entities))
	for _, entity := range entities {
		if parameter := Menu.GetMenusParameters(entity.ID); parameter != nil {
			entity.Parameters = *parameter
		}
		tree[entity.ParentId] = append(tree[entity.ParentId], entity)
	}
	return tree
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取子菜单
func (a *authorityMenu) GetChildrenList(menu *model.AuthorityMenu, treeMap map[string][]model.AuthorityMenu) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		a.GetChildrenList(&menu.Children[i], treeMap)
	}
}
