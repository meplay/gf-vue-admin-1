package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type CasbinRouter struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewCasbinRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &CasbinRouter{router: router, response: &response.Handler{}}
}

func (casbin *CasbinRouter) Init() {
	var router = casbin.router.Group("/casbin")
	{
		router.POST("updateCasbin", casbin.response.Handler()(api.Casbin.Update))
		router.POST("getPolicyPathByAuthorityId", casbin.response.Handler()(api.Casbin.GetPolicyPath))
	}
}
