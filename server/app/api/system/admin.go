package api

import (
	"gf-vue-admin/app/api/system/internal"
	"gf-vue-admin/app/model/system/request"
	"gf-vue-admin/app/service/system"
	"gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Admin = new(admin)

type admin struct{}

// Register 用户注册账号
// @Tags SystemAdmin
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body request.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
func (a *admin) Register(r *ghttp.Request) *response.Response {
	var info request.Register
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "注册失败!"}
	}
	if err := service.Admin.Register(&info); err != nil {
		return &response.Response{Error: err, Message: "注册失败!"}
	}
	return &response.Response{Message: "注册成功!"}
}

// SetUserInfo 设置用户权限
// @Tags SystemAdmin
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateAdmin true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func (a *admin) SetUserInfo(r *ghttp.Request) *response.Response {
	var info request.UpdateAdmin
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "更新用户信息失败!"}
	}
	info.Uuid = internal.Context.GetUserUuid(r)
	if data, err := service.Admin.SetUserInfo(&info); err != nil {
		return &response.Response{Error: err, Message: "更新用户信息失败!"}
	} else {
		return &response.Response{Data: g.Map{"userInfo": data}, Message: "更新用户信息成功!"}
	}
}

// ChangePassword 用户修改密码
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
		return &response.Response{Error: err, Message: "修改密码失败!"}
	}
	info.Uuid = internal.Context.GetUserUuid(r)
	if err := service.Admin.ChangePassword(&info); err != nil {
		return &response.Response{Error: err, Message: "修改密码失败!"}
	}
	return &response.Response{Message: "修改密码成功!"}
}

// SetAuthority 设置用户权限
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
		return &response.Response{Error: err, Message: "设置角色失败!"}
	}
	if err := service.Admin.SetUserAuthority(&info); err != nil {
		return &response.Response{Error: err, Message: "设置角色失败!"}
	}
	return &response.Response{Message: "设置角色失败!"}
}

// Delete 删除用户
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
		return &response.Response{Error: err, Message: "删除失败!"}
	}
	if err := service.Admin.Delete(&info); err != nil {
		return &response.Response{Error: err, Message: "删除失败!"}
	}
	return &response.Response{Message: "删除成功!"}
}

// GetList 分页获取用户列表
// @Tags SystemAdmin
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func (a *admin) GetList(r *ghttp.Request) *response.Response {
	var info request.PageInfo
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "获取列表数据失败!"}
	}
	list, total, err := service.Admin.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, Message: "获取列表数据失败!"}
	}
	return &response.Response{Data: response.NewPageResult(list, total, info.Page, info.PageSize), Message: "获取列表数据成功!"}
}
