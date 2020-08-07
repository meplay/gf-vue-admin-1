package router

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitWorkflowRouter 注册功能api路由
func InitWorkflowRouter() {
	WorkflowRouter := g.Server().Group("workflow").Middleware(
		middleware.JwtAuth,
		middleware.CasbinMiddleware,
	)
	{
		WorkflowRouter.POST("createWorkFlow", v1.CreateWorkFlow) // 创建工作流
	}
}
