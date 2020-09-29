package router

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitJwtRouter 注册邮件相关路由
func InitEmailRouter() {
	EmailRouter := g.Server().Group("email").Middleware(
		middleware.JwtAuth,
		middleware.CasbinMiddleware,
	)
	{
		EmailRouter.POST("emailTest", v1.EmailTest) // 发送测试邮件
	}
}
