package system

import (
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Api = new(api)

type api struct{}

func (a *api) Create(r *ghttp.Request) *response.Response {
	return &response.Response{}
}
