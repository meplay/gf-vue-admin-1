package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"github.com/gogf/gf/net/ghttp"
)

type Admin struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewAdminGroup(router *ghttp.RouterGroup) Router {
	return &Admin{router: router, response: &response.Handler{}}
}

func (a *Admin) Init() {
	var admin = a.router.Group("/user")
	{
		admin.PUT("setUserInfo", a.response.Handler()(api.Admin.Update))             // 设置用户信息
		admin.DELETE("deleteUser", a.response.Handler()(api.Admin.Delete))           // 删除用户
		admin.POST("getUserList", a.response.Handler()(api.Admin.GetList))           // 分页获取用户列表
		admin.POST("changePassword", a.response.Handler()(api.Admin.ChangePassword)) // 修改密码
		admin.POST("setUserAuthority", a.response.Handler()(api.Admin.SetAuthority)) // 设置用户权限
	}
}
