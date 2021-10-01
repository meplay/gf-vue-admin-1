package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*captcha)(nil)

type captcha struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewCaptchaGroup(router *ghttp.RouterGroup) interfaces.Router {
	return &captcha{router: router, response: &response.Handler{}}
}

func (r *captcha) Public() interfaces.Router {
	group := r.router.Group("/base")
	{
		group.POST("captcha", r.response.Handler()(system.Captcha.Captcha))
	}
	return r
}

func (r *captcha) Private() interfaces.Router {
	return r
}

func (r *captcha) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *captcha) PrivateWithoutRecord() interfaces.Router {
	return r
}
