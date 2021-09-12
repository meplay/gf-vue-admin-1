package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var OperationRecord = new(operationRecord)

type operationRecord struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewOperationRecordRouter(router *ghttp.RouterGroup) *operationRecord {
	return &operationRecord{router: router, response: &response.Handler{}}
}

func (m *operationRecord) Private() {
	group := m.router.Group("/menu")
	{
		group.POST("createSysOperationRecord", m.response.Handler()(system.OperationRecord.Create))         // 新建操作日志
		group.GET("findSysOperationRecord", m.response.Handler()(system.OperationRecord.First))             // 根据ID获取操作日志
		group.DELETE("deleteSysOperationRecord", m.response.Handler()(system.OperationRecord.Delete))       // 删除操作日志
		group.DELETE("deleteSysOperationRecordByIds", m.response.Handler()(system.OperationRecord.Deletes)) // 批量删除操作日志
		group.GET("getSysOperationRecordList", m.response.Handler()(system.OperationRecord.GetList))        // 获取操作日志列表
	}
}
