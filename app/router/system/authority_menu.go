package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*authorityMenu)(nil)

type authorityMenu struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewAuthorityMenuRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &authorityMenu{router: router, response: &response.Handler{}}
}

func (r *authorityMenu) Public() interfaces.Router {
	return r
}

func (r *authorityMenu) Private() interfaces.Router {
	return r
}

func (r *authorityMenu) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *authorityMenu) PrivateWithoutRecord() interfaces.Router {
	group := r.router.Group("/menu")
	{
		group.POST("getMenu", r.response.Handler()(system.AuthorityMenu.GetMenu))                   // 获取菜单树
		group.POST("getMenuAuthority", r.response.Handler()(system.AuthorityMenu.GetMenuAuthority)) // 获取指定角色menu
	}
	return r
}
