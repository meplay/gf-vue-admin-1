package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"gf-vue-admin/router/internal"
	"github.com/gogf/gf/net/ghttp"
)

type casbin struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewCasbinRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &casbin{router: router, response: &response.Handler{}}
}

func (c *casbin) Init() {
	group := c.router.Group("/casbin").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("updateCasbin", c.response.Handler()(api.Casbin.Update))
		group.POST("getPolicyPathByAuthorityId", c.response.Handler()(api.Casbin.GetPolicyPath))
	}
}
