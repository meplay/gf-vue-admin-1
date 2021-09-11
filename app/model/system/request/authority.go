package request

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
)

type AuthorityCreate struct {
	ParentId      string `json:"parentId" example:"父角色ID"`
	AuthorityId   string `json:"authorityId" binding:"required" example:"角色ID"`
	AuthorityName string `json:"authorityName" example:"角色名"`
	DefaultRouter string `json:"defaultRouter" example:"登录菜单(默认dashboard)"`
}

func (r *AuthorityCreate) Create() system.Authority {
	return system.Authority{
		ParentId:      r.ParentId,
		AuthorityId:   r.AuthorityId,
		AuthorityName: r.AuthorityName,
		DefaultRouter: r.DefaultRouter,
	}
}

type AuthorityCopy struct {
	Authority      system.Authority `json:"authority" example:"角色信息"`
	OldAuthorityId string           `json:"oldAuthorityId"  example:"旧角色ID"`
}
