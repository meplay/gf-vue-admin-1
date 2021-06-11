package service

import (
	model "flipped-aurora/gf-vue-admin/server/app/model/system"
	"flipped-aurora/gf-vue-admin/server/app/model/system/request"
	"flipped-aurora/gf-vue-admin/server/app/service/system/internal"
	"flipped-aurora/gf-vue-admin/server/library/global"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

// GetMenuTree 获取动态菜单树
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *authorityMenu) GetMenuTree(authorityId string) (menus []model.AuthorityMenu, err error) {
	var menuTree = internal.AuthorityMenu.GetTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		internal.AuthorityMenu.GetChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

// AddMenuAuthority 为角色增加menu树
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authorityMenu) AddMenuAuthority(info *request.AddMenuAuthority) error {
	var entity model.Authority
	if err := global.Db.Where("authority_id = ?", info.AuthorityId).First(&entity).Error; err != nil {
		return err
	}
	entity.Menus = info.Menus
	return internal.Authority().ReplaceMenu(&entity)
}

// GetMenuAuthority 查看当前角色树
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authorityMenu) GetMenuAuthority(info *request.GetAuthorityId) (menus *[]model.AuthorityMenu, err error) {
	entities := make([]model.AuthorityMenu, 0, 10)
	err = global.Db.Where("authority_id = ?", info.AuthorityId).Order("sort").Find(&entities).Error
	return &entities, err
}
