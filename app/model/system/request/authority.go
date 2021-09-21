package request

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
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

func (r *AuthorityCreate) DefaultMenu() *AuthoritySetMenu {
	menus := []system.Menu{{
		Model:     global.Model{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Sort:      1,
		Meta:      system.Meta{Title: "仪表盘", Icon: "setting"},
		Component: "view/dashboard/index.vue",
	}}
	return &AuthoritySetMenu{system.Authority{Menus: menus, AuthorityId: r.AuthorityId}}
}

func (r *AuthorityCreate) DefaultCasbin() []system.Casbin {
	return []system.Casbin{
		{PType: "p", Method: "POST", Path: "/base/login"},
		{PType: "p", Method: "POST", Path: "/menu/getMenu"},
		{PType: "p", Method: "POST", Path: "/user/register"},
		{PType: "p", Method: "PUT", Path: "/user/setUserInfo"},
		{PType: "p", Method: "GET", Path: "/user/getUserInfo"},
		{PType: "p", Method: "POST", Path: "/jwt/jsonInBlacklist"},
		{PType: "p", Method: "POST", Path: "/user/changePassword"},
		{PType: "p", Method: "POST", Path: "/user/setUserAuthority"},
	}
}

type AuthorityCopy struct {
	Authority      system.Authority `json:"authority"`
	OldAuthorityId string           `json:"oldAuthorityId"  example:"旧角色ID"`
}

func (r *AuthorityCopy) ToGetAuthorityId() *common.GetAuthorityId  {
	return &common.GetAuthorityId{AuthorityId: r.OldAuthorityId}
}

func (r *AuthorityCopy) ToCasbinSearch() *CasbinSearch {
	return &CasbinSearch{AuthorityId: r.OldAuthorityId}
}

type AuthorityUpdate struct {
	system.Authority
}

func (r *AuthorityUpdate) Update() system.Authority {
	return r.Authority
}

type AuthorityDelete struct {
	system.Authority
}

type AuthoritySetResources struct {
	system.Authority
}

type AuthoritySetMenu struct {
	system.Authority
}
