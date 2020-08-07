package router

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitCasbinRouter 注册权限相关路由
func InitCasbinRouter() {
	CasbinRouter := g.Server().Group("casbin").Middleware(
		middleware.JwtAuth,
		middleware.CasbinMiddleware,
	)
	{
		CasbinRouter.POST("updateCasbin", v1.UpdateCasbin)
		CasbinRouter.POST("getPolicyPathByAuthorityId", v1.GetPolicyPathByAuthorityId)
		CasbinRouter.GET("casbinTest/:pathParam", v1.CasbinTest)
	}
}
