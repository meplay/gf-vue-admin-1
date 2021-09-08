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

func (u *user) Init() {
	group := u.router.Group("/user")
	{
		group.POST("register", u.response.Handler()(system.User.Register))                     // 用户注册账号
		group.GET("getUserInfo", u.response.Handler()(system.User.GetUserInfo))                // 获取自身信息
		group.PUT("setUserInfo", u.response.Handler()(system.User.SetUserInfo))                // 设置用户信息
		group.POST("setUserAuthority", u.response.Handler()(system.User.SetUserAuthority))     // 设置用户权限
		group.POST("setUserAuthorities", u.response.Handler()(system.User.SetUserAuthorities)) // 设置用户权限组
		group.POST("changePassword", u.response.Handler()(system.User.ChangePassword))         // 用户修改密码
		group.DELETE("deleteUser", u.response.Handler()(system.User.Delete))                   // 删除用户
		group.POST("getUserList", u.response.Handler()(system.User.GetList))                   // 分页获取用户列表
	}
}
