package router

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitAdminsRouter 注册管理员路由
func InitAdminsRouter() {
	UserRouter := g.Server().Group("user").Middleware(
		middleware.JwtAuth,
		middleware.CasbinMiddleware,
	)
	{
		UserRouter.POST("changePassword", v1.ChangePassword)     // 修改密码
		UserRouter.POST("getUserList", v1.GetAdminList)          // 分页获取用户列表
		UserRouter.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("deleteUser", v1.DeleteAdmin)          // 删除用户
		UserRouter.PUT("setUserInfo", v1.SetAdminInfo)           // 设置用户信息
	}
}
