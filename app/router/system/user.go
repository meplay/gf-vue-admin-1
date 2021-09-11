package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type user struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewUserRouter(router *ghttp.RouterGroup) *user {
	return &user{router: router, response: &response.Handler{}}
}

func (r *user) Public() {
	group := r.router.Group("/base")
	{
		group.POST("login", r.response.Handler()(system.User.Login)) // 用户登录
	}
}

func (r *user) Private() {
	group := r.router.Group("/user")
	{
		group.POST("register", r.response.Handler()(system.User.Register))                     // 用户注册账号
		group.GET("getUserInfo", r.response.Handler()(system.User.GetUserInfo))                // 获取自身信息
		group.PUT("setUserInfo", r.response.Handler()(system.User.SetUserInfo))                // 设置用户信息
		group.POST("setUserAuthority", r.response.Handler()(system.User.SetUserAuthority))     // 设置用户权限
		group.POST("setUserAuthorities", r.response.Handler()(system.User.SetUserAuthorities)) // 设置用户权限组
		group.POST("changePassword", r.response.Handler()(system.User.ChangePassword))         // 用户修改密码
		group.DELETE("deleteUser", r.response.Handler()(system.User.Delete))                   // 删除用户
		group.POST("getUserList", r.response.Handler()(system.User.GetList))                   // 分页获取用户列表
	}
}
