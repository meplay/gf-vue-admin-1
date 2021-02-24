package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type base struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewBaseGroup(router *ghttp.RouterGroup) interfaces.Router {
	return &base{router: router, response: &response.Handler{}}
}

func (b *base) Init() {
	group :=  b.router.Group("/base")
	{
		group.POST("captcha", b.response.Handler()(api.Base.Captcha))
		group.POST("login", api.GfJWTMiddleware.LoginHandler) // 登录
	}
}