package router

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitJwtRouter 注册jwt相关路由
func InitAutoCodeRouter() {
	AutoCodeRouter := g.Server().Group("autoCode").Middleware(
		middleware.JwtAuth,
		middleware.CasbinMiddleware,
	)
	{
		AutoCodeRouter.POST("createTemp", v1.CreateTemp) // 创建自动化代码
		AutoCodeRouter.GET("getTables", v1.GetTables)    // 获取对应数据库的表
		AutoCodeRouter.GET("getDB", v1.GetDB)            // 获取数据库
		AutoCodeRouter.GET("getColume", v1.GetColumns)   // 获取指定表所有字段信息
	}
}
