package internal

import (
	model "gf-vue-admin/app/model/system"
	"github.com/gogf/gf/frame/g"
)

var Authority = new(authority)

func (a *authority) Init() {
	entities := make([]model.Authority, 0, 1)
	if err := g.DB().Table(Authority._authority.TableName()).Structs(&entities); err != nil {
		g.Log().Error("获取全部 Authority 失败!", g.Map{"err": err})
	} else {
		Authority.authorityMap = make(map[string]model.Authority, len(entities))
		Authority.authoritiesMap = make(map[string][]model.Authority, len(entities))
		for _, entity := range entities {
			Authority.authorityMap[entity.AuthorityId] = entity
			if entity.ParentId != "0" {
				if value, ok := Authority.authoritiesMap[entity.ParentId]; ok {
					value = append(value, entity)
					Authority.authoritiesMap[entity.ParentId] = value
				} else {
					var a1 = make([]model.Authority, 0, 1)
					a1 = append(a1, entity)
					Authority.authoritiesMap[entity.ParentId] = a1
				}
			}
		}
	}
}

type authority struct {
	_menu             model.Menu
	_authority        model.Authority
	_dataAuthorities  model.DataAuthorities
	_authoritiesMenus model.AuthoritiesMenus

	authorityMap   map[string]model.Authority
	authoritiesMap map[string][]model.Authority
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 查询资源角色
func (a *authority) First(id string) *model.Authority {
	a.Init()
	entity := a.authorityMap[id]
	return &entity
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 查询资源角色
func (a *authority) GetDataAuthority(id string) (result *[]model.Authority) {
	a.Init()
	entities := make([]model.DataAuthorities, 0, 10)
	if err := g.DB().Table(a._dataAuthorities.TableName()).Where(g.Map{"authority_id": id}).Structs(&entities); err != nil {
		g.Log().Error("查询角色的资源角色失败!", g.Map{"err": err})
		return nil
	}
	var authorities = make([]model.Authority, 0, len(entities))
	for _, entity := range entities {
		var a1 = a.authorityMap[entity.DataAuthority]
		authorities = append(authorities, a1)
	}
	return &authorities
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 查询子角色
func (a *authority) FindChildren(authority *model.Authority) {
	a.Init()
	authority.Children = a.authoritiesMap[authority.AuthorityId]
	if len(authority.Children) > 0 {
		for i := range authority.Children {
			a.FindChildren(&authority.Children[i])
		}
	}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据 authority_id 获取menu列表数据
func (a *authority) GetMenus(id string) *[]model.Menu {
	entities := make([]model.AuthoritiesMenus, 0, 10)
	if err := g.DB().Table(a._authoritiesMenus.TableName()).Where(g.Map{"authority_id": id}).Structs(&entities); err != nil {
		g.Log().Error("获取 authorities_menus 表数据失败!", g.Map{"authority_id": id})
		return nil
	}
	menus := make([]model.Menu, 0, 10)
	for _, entity := range entities {
		var m1 model.Menu
		if err := g.DB().Table(a._menu.TableName()).WherePri(entity.MenuId).Struct(&m1); err != nil {
			return &menus
		} else {
			menus = append(menus, m1)
		}
	}
	return &menus
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除原有menu树
func (a *authority) ReplaceMenu(info *model.Authority) error {
	if _, err := g.DB().Table(a._authoritiesMenus.TableName()).Delete(g.Map{"authority_id": info.AuthorityId}); err != nil {
		g.Log().Error("删除 authorities_menus 表数据失败!", g.Map{"authority_id": info.AuthorityId, "err": err})
		return err
	}
	for _, m := range info.Menus {
		if _, err := g.DB().Table(a._authoritiesMenus.TableName()).Insert(&model.AuthoritiesMenus{MenuId: m.ID, AuthorityId: info.AuthorityId}); err != nil {
			return err
		}
	}
	return nil
}
