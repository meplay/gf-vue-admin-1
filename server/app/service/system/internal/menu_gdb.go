package internal

import (
	model "gf-vue-admin/app/model/system"
	"github.com/gogf/gf/frame/g"
	"strconv"
)

var Menu = new(menu)

type menu struct {
	_menu             model.Menu
	_parameters       model.MenuParameter
	_menusParameters  model.MenusParameters
	_authoritiesMenus model.AuthoritiesMenus
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取菜单的Parameters
func (m *menu) GetMenusParameters(id uint) *[]model.MenuParameter {
	entities := make([]model.MenusParameters, 0, 10)
	if err := g.DB().Table(m._menusParameters.TableName()).Where(g.Map{"menu_id": id}).Structs(&entities); err != nil {
		g.Log().Error("获取 menus_parameters 表数据失败!", g.Map{"menu_id": id})
		return nil
	}
	parameters := make([]model.MenuParameter, 0, len(entities))
	for _, entity := range entities {
		var e1 model.MenuParameter
		if err := g.DB().Table(m._parameters.TableName()).Where(g.Map{"parameter_id": entity.ParameterId}).Struct(&e1); err != nil {
			g.Log().Error("获取 menus_parameters 表数据失败", g.Map{"parameter_id": entity.ParameterId})
		} else {
			parameters = append(parameters, e1)
		}
	}
	return &parameters
}

func (m *menu) GetAuthoritiesMenus(id uint) *[]model.Authority {
	Authority.Init()
	entities := make([]model.AuthoritiesMenus, 0, 10)
	if err := g.DB().Table(m._menusParameters.TableName()).Where(g.Map{"menu_id": id}).Structs(&entities); err != nil {
		g.Log().Error("获取 authorities_menus 表数据失败!", g.Map{"menu_id": id})
		return nil
	}
	authorities := make([]model.Authority, 0, len(entities))
	for _, entity := range entities {
		authorities = append(authorities, Authority.authorityMap[entity.AuthorityId])
	}
	return &authorities
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取菜单的子菜单
func (m *menu) GetChildrenList(menu *model.Menu, treeMap map[string][]model.Menu) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		m.GetChildrenList(&menu.Children[i], treeMap)
	}
	return
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取路由总树map
func (m *menu) GetTreeMap() map[string][]model.Menu {
	var entities []model.Menu
	if err := g.DB().Table(m._menu.TableName()).Structs(&entities); err != nil {
		g.Log().Error("获取 menus 表数据失败!")
		return nil
	}
	tree := make(map[string][]model.Menu, len(entities))
	for _, entity := range entities {
		entity.Parameters = *Menu.GetMenusParameters(entity.ID)
		tree[entity.ParentId] = append(tree[entity.ParentId], entity)
	}
	return tree
}
