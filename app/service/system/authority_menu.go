package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

// GetAuthorityMenu 查看当前角色树
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authorityMenu) GetAuthorityMenu(info *common.GetAuthorityId) (menus []system.AuthorityMenu, err error) {
	err = global.Db.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	return menus, err
}
