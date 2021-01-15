package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"server/app/api/internal"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/service/system"
	"server/library/global"
)

var Admin = new(admin)

type admin struct{}

// @Tags SystemAdmin
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePassword true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func (a *admin) ChangePassword(r *ghttp.Request) *response.Response {
	var info request.ChangePassword
	if err := r.Parse(&info); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	info.Uuid = internal.Info.GetUserUuid(r)
	if err := service.Admin.ChangePassword(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorChangePassword}
	}
	return &response.Response{MessageCode: response.SuccessChangePassword}
}

// @Tags SystemAdmin
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func (a *admin) GetAdminList(r *ghttp.Request) *response.Response {
	var info request.PageInfo
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.Admin.GetAdminList(&info)
	if err != nil {
		return &response.Response{MessageCode: response.ErrorGetList, Error: err}
	}
	return &response.Response{MessageCode: response.SuccessGetList, Error: err, Data: response.PageResult{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}}
}

// @Tags SystemAdmin
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetAuthority true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func (a *admin) SetAuthority(r *ghttp.Request) *response.Response {
	var info request.SetAuthority
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorSetAuthority}
	}
	if err := service.Admin.SetUserAuthority(&info); err != nil {
		return &response.Response{MessageCode: response.ErrorSetAuthority, Error: err}
	}
	return &response.Response{MessageCode: response.SuccessSetAuthority}
}

// @Tags SystemAdmin
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
func (a *admin) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.Admin.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags SystemAdmin
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateAdmin true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func (a *admin) Update(r *ghttp.Request) *response.Response {
	var info request.UpdateAdmin
	if err := r.Parse(&info); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	info.Uuid = internal.Info.GetUserUuid(r)
	if data, err := service.Admin.SetAdminInfo(&info); err != nil {
		return &response.Response{MessageCode: response.ErrorSetAdminInfo, Error: err}
	} else {
		return &response.Response{MessageCode: response.SuccessSetAdminInfo, Data: g.Map{"userInfo": data}}
	}
}
