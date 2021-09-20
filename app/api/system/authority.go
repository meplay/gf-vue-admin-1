package system

import (
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Authority = new(authority)

type authority struct{}

func (a *authority) Create(r *ghttp.Request) *response.Response {
	return &response.Response{}
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