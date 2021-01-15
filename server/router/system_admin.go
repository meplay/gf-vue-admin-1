package router

import (
	"github.com/gogf/gf/net/ghttp"
	"server/app/api/response"
	api "server/app/api/system"
)

type AdminRouter struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewAdminRouter(router *ghttp.RouterGroup) *AdminRouter {
	return &AdminRouter{router: router, response: &response.Handler{}}
}

func (a *AdminRouter) Init() {
	var admin = a.router.Group("user")
	{
		admin.PUT("setUserInfo", a.response.Handler()(api.Admin.Update))             // 设置用户信息
		admin.POST("getUserList", a.response.Handler()(api.Admin.GetAdminList))      // 分页获取用户列表
		admin.DELETE("deleteUser", a.response.Handler()(api.Admin.Delete))           // 删除用户
		admin.POST("changePassword", a.response.Handler()(api.Admin.ChangePassword)) // 修改密码
		admin.POST("setUserAuthority", a.response.Handler()(api.Admin.SetAuthority)) // 设置用户权限
	}
}
