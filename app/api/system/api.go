package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Api = new(api)

type api struct{}

// Create
// @Tags SystemApi
// @Summary 创建api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ApiCreate true "请求参数"
// @Success 200 {object} response.Response{} "创建成功!"
// @Router /api/createApi [post]
func (a *api) Create(r *ghttp.Request) *response.Response {
	var info request.ApiCreate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := system.Api.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// First
// @Tags SystemApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "获取数据成功!"
// @Router /api/getApiById [post]
func (a *api) First(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	data, err := system.Api.First(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	return &response.Response{Data: g.Map{"api": data}, MessageCode: response.SuccessFirst}
}

// Update
// @Tags SystemApi
// @Summary 更新api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ApiUpdate true "请求参数"
// @Success 200 {object} response.Response{data=request.ApiUpdate} "更新成功!"
// @Router /api/updateApi [post]
func (a *api) Update(r *ghttp.Request) *response.Response {
	var info request.ApiUpdate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := system.Api.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	return &response.Response{Data: g.Map{"api": info}, MessageCode: response.SuccessUpdated}
}

// Delete
// @Tags SystemApi
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeleteApi true "请求参数"
// @Success 200 {object} response.Response{} "删除成功!"
// @Router /api/deleteApi [post]
func (a *api) Delete(r *ghttp.Request) *response.Response {
	var info request.DeleteApi
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := system.Api.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// Deletes
// @Tags SystemApi
// @Summary 批量删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByIDs true "请求参数"
// @Success 200 {object} response.Response{} "批量删除成功!"
// @Router /api/deleteApisByIds [delete]
func (a *api) Deletes(r *ghttp.Request) *response.Response {
	var info common.GetByIDs
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := system.Api.Deletes(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// GetList
// @Tags SystemApi
// @Summary 分页获取api列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ApiSearch true "请求参数"
// @Success 200 {object} response.Response{data=[]system.Api} "获取列表数据成功!"
// @Router /api/getApiList [post]
func (a *api) GetList(r *ghttp.Request) *response.Response {
	var info request.ApiSearch
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := system.Api.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: common.NewPageResult(list, total, info.PageInfo), MessageCode: response.SuccessGetList}
}

// GetAllApis
// @Tags SystemApi
// @Summary 获取所有的Api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]system.Api} "获取列表数据成功!"
// @Router /api/getAllApis [post]
func (a *api) GetAllApis(r *ghttp.Request) *response.Response {
	if apis, err := system.Api.GetAllApis(); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	} else {
		return &response.Response{Data: g.Map{"apis": apis}, MessageCode: response.SuccessUpdated}
	}
}
