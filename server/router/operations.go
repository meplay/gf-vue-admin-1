package router

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitOperationRouter 注册操作记录路由
func InitOperationRouter() {
	OperationRouter := g.Server().Group("sysOperationRecord").Middleware(
		middleware.JwtAuth,
		middleware.CasbinMiddleware,
	)
	{
		OperationRouter.POST("createOperationRecord", v1.CreateOperation)            // 新建Operation
		OperationRouter.DELETE("deleteSysOperationRecordById", v1.DeleteOperation)   // 删除Operation
		OperationRouter.DELETE("deleteSysOperationRecordByIds", v1.DeleteOperations) // 批量删除Operation
		OperationRouter.PUT("updateOperationRecord", v1.UpdateOperation)             // 更新Operation
		OperationRouter.GET("findOperationRecord", v1.FindOperation)                 // 根据ID获取Operation
		OperationRouter.GET("getSysOperationRecordList", v1.GetOperationList)        // 获取Operation列表
	}
}
