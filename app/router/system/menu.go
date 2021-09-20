package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/app/router/middleware"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*menu)(nil)

type menu struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewMenuRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &menu{router: router, response: &response.Handler{}}
}

func (r *menu) Public() interfaces.Router {
	return r
}

func (r *menu) Private() interfaces.Router {
	group := r.router.Group("/menu").Middleware(middleware.OperationRecord)
	{
		group.POST("addBaseMenu", r.response.Handler()(system.Menu.Create))                // 新增菜单
		group.POST("updateBaseMenu", r.response.Handler()(system.Menu.Update))             // 更新菜单
		group.POST("deleteBaseMenu", r.response.Handler()(system.Menu.Delete))             // 删除菜单
		group.POST("addMenuAuthority", r.response.Handler()(system.Menu.AddMenuAuthority)) // 增加menu和角色关联关系
	}
	return r
}

func (r *menu) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *menu) PrivateWithoutRecord() interfaces.Router {
	group := r.router.Group("/menu")
	{
		group.POST("getBaseMenuById", r.response.Handler()(system.Menu.First))             // 根据id获取菜单
		group.POST("getMenuList", r.response.Handler()(system.Menu.GetList))               // 分页获取基础menu列表
		group.POST("getBaseMenuTree", r.response.Handler()(system.Menu.GetTree))           // 获取用户动态路由

	}
	return r
}