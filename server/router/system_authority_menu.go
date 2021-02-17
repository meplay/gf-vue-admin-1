package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type MenuRouter struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewMenuRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &MenuRouter{router: router, response: &response.Handler{}}
}

func (m *MenuRouter) Init() {
	var menu = m.router.Group("/menu")
	{
		menu.POST("addBaseMenu", m.response.Handler()(api.Menu.Create))                         // 新增菜单
		menu.POST("getBaseMenuById", m.response.Handler()(api.Menu.First))                      // 根据id获取菜单
		menu.POST("updateBaseMenu", m.response.Handler()(api.Menu.Update))                      // 更新菜单
		menu.POST("deleteBaseMenu", m.response.Handler()(api.Menu.Delete))                      // 删除菜单
		menu.POST("getMenuList", m.response.Handler()(api.Menu.GetList))                        // 分页获取基础menu列表
		menu.POST("getBaseMenuTree", m.response.Handler()(api.Menu.GetBaseMenuTree))            // 获取用户动态路由
		menu.POST("getMenu", m.response.Handler()(api.AuthorityMenu.GetMenu))                   // 获取菜单树
		menu.POST("addMenuAuthority", m.response.Handler()(api.AuthorityMenu.AddMenuAuthority)) // 增加menu和角色关联关系
		menu.POST("getMenuAuthority", m.response.Handler()(api.AuthorityMenu.GetMenuAuthority)) // 获取指定角色menu
	}
}
