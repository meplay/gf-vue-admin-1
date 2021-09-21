package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Casbin = new(casbin)

type casbin struct{}

// Update
// @Tags SystemCasbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "请求参数"
// @Success 200 {object} response.Response{} "更新成功!"
// @Router /casbin/UpdateCasbin [post]
func (a *casbin) Update(r *ghttp.Request) *response.Response {
	var info request.CasbinInReceive
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := system.Casbin.Update(info.AuthorityId, info.CasbinInfos); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	return &response.Response{MessageCode: response.SuccessUpdated}
}

// GetPolicyPathByAuthorityId
// @Tags SystemCasbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "请求参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func (a *casbin) GetPolicyPathByAuthorityId(r *ghttp.Request) *response.Response {
	var info request.CasbinInReceive
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "获取失败!"}
	}
	paths := system.Casbin.GetPolicyPathByAuthorityId(info.AuthorityId)
	return &response.Response{Data: g.Map{"paths": paths}, Message: "获取成功!"}
}
