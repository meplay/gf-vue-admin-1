package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var DictionaryDetail = new(dictionaryDetail)

type dictionaryDetail struct{}

// Create
// @Tags SystemDictionaryDetail
// @Summary 创建系统字典详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionaryDetail true "请求参数"
// @Success 200 {object} response.Response{message=string} "创建成功!"
// @Router /sysDictionaryDetail/createSysDictionaryDetail [post]
func (a *dictionaryDetail) Create(r *ghttp.Request) *response.Response {
	var info request.DictionaryDetailCreate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := system.DictionaryDetail.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// First
// @Tags SystemDictionaryDetail
// @Summary 用id查询系统详情字典
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysDictionaryDetail true "请求参数"
// @Success 200 {object} response.Response{message=string} "获取数据成功!"
// @Router /sysDictionaryDetail/findSysDictionaryDetail [get]
func (a *dictionaryDetail) First(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	data, err := system.DictionaryDetail.First(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	return &response.Response{Data: g.Map{"resysDictionaryDetail": data}, MessageCode: response.SuccessFirst}
}

// Update
// @Tags SystemDictionaryDetail
// @Summary 更新系统详情字典
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionaryDetail true "请求参数"
// @Success 200 {object} response.Response{data=system.DictionaryDetail} "更新成功!"
// @Router /sysDictionaryDetail/updateSysDictionaryDetail [put]
func (a *dictionaryDetail) Update(r *ghttp.Request) *response.Response {
	var info request.DictionaryDetailUpdate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := system.DictionaryDetail.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	return &response.Response{MessageCode: response.SuccessUpdated}
}

// Delete
// @Tags SystemDictionaryDetail
// @Summary 删除系统详情字典
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionaryDetail true "请求参数"
// @Success 200 {object} response.Response{message=string} "删除成功!"
// @Router /sysDictionaryDetail/deleteSysDictionaryDetail [delete]
func (a *dictionaryDetail) Delete(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := system.DictionaryDetail.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// GetList
// @Tags SystemDictionaryDetail
// @Summary 分页获取系统字典详情列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysDictionaryDetailSearch true "请求参数"
// @Success 200 {object} response.Response{data=[]system.DictionaryDetail} "获取列表数据成功!"
// @Router /sysDictionaryDetail/getSysDictionaryDetailList [get]
func (a *dictionaryDetail) GetList(r *ghttp.Request) *response.Response {
	var info request.DictionaryDetailSearch
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := system.DictionaryDetail.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: common.NewPageResult(list, total, info.PageInfo), MessageCode: response.SuccessGetList}
}
