package system

import (
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Base = new(base)

type base struct{}

func (b *base) Captcha(r *ghttp.Request) *response.Response {
	return &response.Response{}
}
