package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type authorityMenu struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewAuthorityMenuRouter(router *ghttp.RouterGroup) *authorityMenu {
	return &authorityMenu{router: router, response: &response.Handler{}}
}

func (m *authorityMenu) Private() {
	group := m.router.Group("/menu")
	{
		group.POST("getMenu", m.response.Handler()(system.AuthorityMenu.GetMenu))                   // 获取菜单树
		group.POST("getMenuAuthority", m.response.Handler()(system.AuthorityMenu.GetMenuAuthority)) // 获取指定角色menu

	}
}
