package router

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitJwtRouter 注册jwt相关路由
func InitJwtRouter() {
	ApiRouter := g.Server().Group("jwt").Middleware(
		middleware.JwtAuth,
		middleware.CasbinMiddleware,
	)
	{
		ApiRouter.POST("jsonInBlacklist", v1.JsonInBlacklist) // jwt加入黑名单
	}
}
