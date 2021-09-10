package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type captcha struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewCaptchaGroup(router *ghttp.RouterGroup) *captcha {
	return &captcha{router: router, response: &response.Handler{}}
}

func (r *captcha) Public() {
	group := r.router.Group("/base")
	{
		group.POST("captcha", r.response.Handler()(system.Captcha.Captcha))
	}
}
