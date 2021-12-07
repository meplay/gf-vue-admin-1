package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var AutoCodeHistory = new(autoCodeHistory)

type autoCodeHistory struct{}

// First
// @Tags SystemAutoCodeHistory
// @Summary 回滚
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "获取成功!"
// @Router /autoCode/getMeta [post]
func (a *autoCodeHistory) First(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	data, err := system.AutoCodeHistory.First(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	return &response.Response{Data: g.Map{"meta": data}, MessageCode: response.ErrorFirst}
}

// Delete
// @Tags AutoCode
// @Summary 删除回滚记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "删除成功!"
// @Router /autoCode/delSysHistory [post]
func (a *autoCodeHistory) Delete(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := system.AutoCodeHistory.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// RollBack
// @Tags SystemAutoCodeHistory
// @Summary 回滚
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "回滚成功!"
// @Router /autoCode/rollback [post]
func (a *autoCodeHistory) RollBack(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "回滚失败!"}
	}
	if err := system.AutoCodeHistory.RollBack(&info); err != nil {
		return &response.Response{Error: err, Message: "回滚失败!"}
	}
	return &response.Response{Message: "回滚成功!"}
}

// GetList
// @Tags SystemAutoCodeHistory
// @Summary 查询回滚记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AutoCodeHistorySearch true "请求参数"
// @Success 200 {object} response.Response{} "回滚成功!"
// @Router /autoCode/getSysHistory [post]
func (a *autoCodeHistory) GetList(r *ghttp.Request) *response.Response {
	var info request.AutoCodeHistorySearch
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := system.AutoCodeHistory.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: common.NewPageResult(list, total, *info.PageInfo), MessageCode: response.SuccessGetList}
}
