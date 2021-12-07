package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var OperationRecord = new(operationRecord)

type operationRecord struct{}

// Create
// @Tags SystemOperationRecord
// @Summary 创建操作日志记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.OperationRecordCreate true "请求参数"
// @Success 200 {object} response.Response{} "创建成功!"
// @Router /sysOperationRecord/createSysOperationRecord [post]
func (a *operationRecord) Create(r *ghttp.Request) *response.Response {
	var info request.OperationRecordCreate
	if err := system.OperationRecord.Create(&info); err != nil {
		return &response.Response{MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// First
// @Tags SystemOperationRecord
// @Summary 用id查询操作日志记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "获取数据成功!"
// @Router /sysOperationRecord/findSysOperationRecord [get]
func (a *operationRecord) First(r *ghttp.Request) *response.Response {
	var info common.GetByID
	data, err := system.OperationRecord.First(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	return &response.Response{Data: g.Map{"resysOperationRecord": data}, MessageCode: response.SuccessFirst}
}

// Delete
// @Tags SystemOperationRecord
// @Summary 删除操作日志记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "删除成功!"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
func (a *operationRecord) Delete(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := system.OperationRecord.Delete(&info); err != nil {
		return &response.Response{MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// Deletes
// @Tags SystemOperationRecord
// @Summary 批量删除操作日志记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "批量删除成功!"
// @Router /sysOperationRecord/deleteSysOperationRecordByIds [delete]
func (a *operationRecord) Deletes(r *ghttp.Request) *response.Response {
	var info common.GetByIDs
	if err := system.OperationRecord.Deletes(&info); err != nil {
		return &response.Response{MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// GetList
// @Tags SystemOperationRecord
// @Summary 分页获取操作日志记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.OperationRecordSearch true "请求参数"
// @Success 200 {object} response.Response{data=[]system.OperationRecord} "获取列表数据成功!"
// @Router /sysOperationRecord/getSysOperationRecordList [get]
func (a *operationRecord) GetList(r *ghttp.Request) *response.Response {
	var info request.OperationRecordSearch
	list, total, err := system.OperationRecord.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: common.NewPageResult(list, total, info.PageInfo), MessageCode: response.SuccessGetList}
}
