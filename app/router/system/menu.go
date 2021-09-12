package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type menu struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewMenuRouter(router *ghttp.RouterGroup) *menu {
	return &menu{router: router, response: &response.Handler{}}
}

func (m *menu) Private() {
	group := m.router.Group("/menu")
	{
		group.POST("addBaseMenu", m.response.Handler()(system.Menu.Create))                // 新增菜单
		group.POST("getBaseMenuById", m.response.Handler()(system.Menu.First))             // 根据id获取菜单
		group.POST("updateBaseMenu", m.response.Handler()(system.Menu.Update))             // 更新菜单
		group.POST("deleteBaseMenu", m.response.Handler()(system.Menu.Delete))             // 删除菜单
		group.POST("getMenuList", m.response.Handler()(system.Menu.GetList))               // 分页获取基础menu列表
		group.POST("getBaseMenuTree", m.response.Handler()(system.Menu.GetTree))           // 获取用户动态路由
		group.POST("addMenuAuthority", m.response.Handler()(system.Menu.AddMenuAuthority)) // 增加menu和角色关联关系
	}
}
