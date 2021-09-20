package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/app/router/middleware"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*user)(nil)

type user struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewUserRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &user{router: router, response: &response.Handler{}}
}

func (r *user) Public() interfaces.Router {
	group := r.router.Group("/base")
	{
		group.POST("login", r.response.Handler()(system.User.Login)) // 用户登录
	}
	return r
}

func (r *user) Private() interfaces.Router {
	group := r.router.Group("/user").Middleware(middleware.OperationRecord)
	{
		group.POST("register", r.response.Handler()(system.User.Register))                     // 用户注册账号
		group.PUT("setUserInfo", r.response.Handler()(system.User.SetUserInfo))                // 设置用户信息
		group.POST("setUserAuthority", r.response.Handler()(system.User.SetUserAuthority))     // 设置用户权限
		group.POST("setUserAuthorities", r.response.Handler()(system.User.SetUserAuthorities)) // 设置用户权限组
		group.POST("changePassword", r.response.Handler()(system.User.ChangePassword))         // 用户修改密码
		group.DELETE("deleteUser", r.response.Handler()(system.User.Delete))                   // 删除用户
	}
	return r
}

func (r *user) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *user) PrivateWithoutRecord() interfaces.Router {
	group := r.router.Group("/user")
	{
		group.GET("getUserInfo", r.response.Handler()(system.User.GetUserInfo)) // 获取自身信息
		group.POST("getUserList", r.response.Handler()(system.User.GetList))    // 分页获取用户列表
	}
	return r
}
