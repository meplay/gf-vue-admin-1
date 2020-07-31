package router

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitCasbinRuleRouter 注册casbinRule路由
func InitCasbinRuleRouter() {
	CasbinRuleRouter := g.Server().Group("casbinRule").Middleware(middleware.JwtAuth).Middleware(middleware.CasbinMiddleware)
	{
		CasbinRuleRouter.POST("createCasbinRule", v1.CreateCasbinRule)               // 新建CasbinRule
		CasbinRuleRouter.DELETE("deleteCasbinRule", v1.DeleteCasbinRule)             // 删除CasbinRule
        CasbinRuleRouter.DELETE("deleteCasbinRuleByIds", v1.DeleteCasbinRuleByIds)   // 批量删除CasbinRule
		CasbinRuleRouter.PUT("updateCasbinRule", v1.UpdateCasbinRule)                // 更新CasbinRule
		CasbinRuleRouter.GET("findCasbinRule", v1.FindCasbinRule)                    // 根据ID获取CasbinRule
		CasbinRuleRouter.GET("getCasbinRuleList", v1.GetCasbinRuleList)              // 获取CasbinRule列表
	}
}
