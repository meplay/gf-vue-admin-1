package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Authority = new(authority)

type authority struct{}

// Create
// @Tags SystemAuthority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AuthorityCreate true "请求参数"
// @Success 200 {object} response.Response{data=system.Authority} "创建成功!"
// @Router /authority/createAuthority [post]
func (a *authority) Create(r *ghttp.Request) *response.Response {
	var info request.AuthorityCreate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	data, err := system.Authority.Create(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{Data: g.Map{"authority": data}, MessageCode: response.SuccessCreated}
}

// Copy
// @Tags SystemAuthority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AuthorityCopy true "请求参数"
// @Success 200 {object} response.Response{data=system.Authority} "拷贝成功!"
// @Router /authority/copyAuthority [post]
func (a *authority) Copy(r *ghttp.Request) *response.Response {
	var info request.AuthorityCopy
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, Message: "拷贝角色失败!"}
	}
	data, err := system.Authority.Copy(&info)
	if err != nil {
		return &response.Response{Error: err, Message: "拷贝角色失败!"}
	}
	return &response.Response{Data: g.Map{"authority": data}, Message: "拷贝成功!"}
}

// Update
// @Tags SystemAuthority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AuthorityUpdate true "请求参数"
// @Success 200 {object} response.Response{data=system.Authority} "更新成功!"
// @Router /authority/updateAuthority [post]
func (a *authority) Update(r *ghttp.Request) *response.Response {
	var info request.AuthorityUpdate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	data, err := system.Authority.Update(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	return &response.Response{Data: g.Map{"authority": data}, MessageCode: response.SuccessUpdated}
}

// Delete
// @Tags SystemAuthority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AuthorityDelete true "请求参数"
// @Success 200 {object} response.Response{} "删除成功!"
// @Router /authority/deleteAuthority [post]
func (a *authority) Delete(r *ghttp.Request) *response.Response {
	var info request.AuthorityDelete
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := system.Authority.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// GetList
// @Tags SystemAuthority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.PageInfo true "请求参数"
// @Success 200 {object} response.Response{data=[]system.Authority} "获取列表数据成功!"
// @Router /authority/getAuthorityList [post]
func (a *authority) GetList(r *ghttp.Request) *response.Response {
	var info common.PageInfo
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := system.Authority.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: common.NewPageResult(list, total, info), MessageCode: response.SuccessGetList}
}

func (a *authority) SetAuthorityResources(r *ghttp.Request) *response.Response {
	return &response.Response{}
}
