package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*operationRecord)(nil)

type operationRecord struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewOperationRecordRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &operationRecord{router: router, response: &response.Handler{}}
}

func (r *operationRecord) Public() interfaces.Router {
	return r
}

func (r *operationRecord) Private() interfaces.Router {
	group := r.router.Group("/sysOperationRecord")
	{
		group.POST("createSysOperationRecord", r.response.Handler()(system.OperationRecord.Create))         // 新建操作日志
		group.DELETE("deleteSysOperationRecord", r.response.Handler()(system.OperationRecord.Delete))       // 删除操作日志
		group.DELETE("deleteSysOperationRecordByIds", r.response.Handler()(system.OperationRecord.Deletes)) // 批量删除操作日志
	}
	return r
}

func (r *operationRecord) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *operationRecord) PrivateWithoutRecord() interfaces.Router {
	group := r.router.Group("/sysOperationRecord")
	{
		group.GET("findSysOperationRecord", r.response.Handler()(system.OperationRecord.First))      // 根据ID获取操作日志
		group.GET("getSysOperationRecordList", r.response.Handler()(system.OperationRecord.GetList)) // 获取操作日志列表
	}
	return r
}
