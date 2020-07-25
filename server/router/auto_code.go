package router

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitJwtRouter 注册jwt相关路由
func InitAutoCodeRouter() {
	AutoCodeRouter := g.Server().Group("autoCode").Middleware(middleware.JwtAuth).Middleware(middleware.CasbinMiddleware)
	{
		AutoCodeRouter.POST("createTemp", v1.CreateTemp) // 创建自动化代码
	}
}
