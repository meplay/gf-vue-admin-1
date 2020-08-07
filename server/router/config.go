package router

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitSystemRouter 注册system配置相关路由
func InitSystemRouter() {
	ConfigRouter := g.Server().Group("system").Middleware(
		middleware.JwtAuth,
		middleware.CasbinMiddleware,
	)
	{
		ConfigRouter.POST("getSystemConfig", v1.GetSystemConfig) // 获取配置文件内容
		ConfigRouter.POST("setSystemConfig", v1.SetSystemConfig) // 设置配置文件内容
	}
}
