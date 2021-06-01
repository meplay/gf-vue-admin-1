package internal

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
)

func Authority() *authority {
	var _authority authority
	entities := make([]model.Authority, 0, 1)
	if err := global.Db.Find(&entities).Error;err != nil {
		g.Log().Error("获取全部 Authority 失败!", g.Map{"err": err})
	}else {
		_authority.authorityMap = make(map[string]model.Authority, len(entities))
		_authority.authoritiesMap = make(map[string][]model.Authority, len(entities))
		for _, entity := range entities {
			_authority.authorityMap[entity.AuthorityId] = entity
			if entity.ParentId != "0" {
				if value, ok := _authority.authoritiesMap[entity.ParentId]; ok {
					value = append(value, entity)
					_authority.authoritiesMap[entity.ParentId] = value
				} else {
					var a1 = make([]model.Authority, 0, 1)
					a1 = append(a1, entity)
					_authority.authoritiesMap[entity.ParentId] = a1
				}
			}
		}
	}
	return &_authority
}

type authority struct {
	_menu             model.Menu
	_authority        model.Authority
	_dataAuthorities  model.DataAuthorities
	_authoritiesMenus model.AuthoritiesMenus

	authorityMap   map[string]model.Authority
	authoritiesMap map[string][]model.Authority
}

// First 查询资源角色
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *authority) First(id string) model.Authority {
	entity := a.authorityMap[id]
	return entity
}

// GetMenus 根据 authority_id 获取menu列表数据
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authority) GetMenus(id string) []model.Menu {
	var entities []model.AuthoritiesMenus
	if err := global.Db.Where("authority_id = ?", id).Find(&entities).Error; err != nil {
		g.Log().Error("获取 authorities_menus 表数据失败!", g.Map{"authority_id": id})
		return nil
	}
	menus := make([]model.Menu, 0, len(entities))
	for _, entity := range entities {
		var m1 model.Menu
		if err := global.Db.Where("id = ?", entity.MenuId).First(&m1).Error; err != nil {
			return menus
		} else {
			menus = append(menus, m1)
		}
	}
	return menus
}

// ReplaceMenu 删除原有menu树
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authority) ReplaceMenu(info *model.Authority) error {
	if err := global.Db.Delete(&model.AuthoritiesMenus{}, info.AuthorityId).Error; err != nil {
		g.Log().Error("删除 authorities_menus 表数据失败!", g.Map{"authority_id": info.AuthorityId, "err": err})
		return err
	}
	for _, m := range info.Menus {
		entity := &model.AuthoritiesMenus{MenuId: m.ID, AuthorityId: info.AuthorityId}
		if err := global.Db.Create(&entity).Error; err != nil {
			return err
		}
	}
	return nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 查询子角色
func (a *authority) FindChildren(authority *model.Authority) {
	authority.Children = a.authoritiesMap[authority.AuthorityId]
	if len(authority.Children) > 0 {
		for i := range authority.Children {
			a.FindChildren(&authority.Children[i])
		}
	}
}

// GetDataAuthority 查询资源角色
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authority) GetDataAuthority(id string) []model.Authority {
	entities := make([]model.DataAuthorities, 0, 10)
	if err := global.Db.Where("authority_id", id).Find(&entities).Error; err != nil {
		g.Log().Error("查询角色的资源角色失败!", g.Map{"err": err})
		return nil
	}
	var authorities = make([]model.Authority, 0, len(entities))
	for _, entity := range entities {
		var a1 = a.authorityMap[entity.DataAuthority]
		authorities = append(authorities, a1)
	}
	return authorities
}
