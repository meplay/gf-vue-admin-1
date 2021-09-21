package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
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
// @Success 200 {object} response.Response{data=[]system.Authority} "获取列表数据成功!"
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

func (a *authority) Copy(r *ghttp.Request) *response.Response {
	return &response.Response{}
}

func (a *authority) Update(r *ghttp.Request) *response.Response {
	return &response.Response{}
}

func (a *authority) Delete(r *ghttp.Request) *response.Response {
	return &response.Response{}
}

func (a *authority) GetList(r *ghttp.Request) *response.Response {
	return &response.Response{}
}

func (a *authority) SetAuthorityResources(r *ghttp.Request) *response.Response {
	return &response.Response{}
}
