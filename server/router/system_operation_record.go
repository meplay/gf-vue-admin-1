package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"github.com/gogf/gf/net/ghttp"
)

type OperationRecordRouter struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewOperationRecordRouter(router *ghttp.RouterGroup) *OperationRecordRouter {
	return &OperationRecordRouter{router: router, response: &response.Handler{}}
}

func (o *OperationRecordRouter) Init() {
	record := o.router.Group("/sysOperationRecord")
	{
		record.POST("createSysOperationRecord", o.response.Handler()(api.OperationRecord.Create))         // 新建SysOperationRecord
		record.GET("findSysOperationRecord", o.response.Handler()(api.OperationRecord.First))             // 根据ID获取SysOperationRecord
		record.DELETE("deleteSysOperationRecord", o.response.Handler()(api.OperationRecord.Delete))       // 删除SysOperationRecord
		record.DELETE("deleteSysOperationRecordByIds", o.response.Handler()(api.OperationRecord.Deletes)) // 批量删除SysOperationRecord
		record.GET("getSysOperationRecordList", o.response.Handler()(api.OperationRecord.GetList))        // 获取SysOperationRecord列表
	}
}
