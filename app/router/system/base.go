package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type base struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewBaseGroup(router *ghttp.RouterGroup) *base {
	return &base{router: router, response: &response.Handler{}}
}

func (b *base) Init() {
	group := b.router.Group("/base")
	{
		group.POST("captcha", b.response.Handler()(system.Base.Captcha))
		group.POST("login", b.response.Handler()(system.Base.Login))
	}
}
