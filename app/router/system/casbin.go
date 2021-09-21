package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/app/router/middleware"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*casbin)(nil)

type casbin struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewCasbinRouter(router *ghttp.RouterGroup) *casbin {
	return &casbin{router: router, response: &response.Handler{}}
}

func (r *casbin) Public() interfaces.Router {
	return r
}

func (r *casbin) Private() interfaces.Router {
	group := r.router.Group("/casbin").Middleware(middleware.OperationRecord)
	{
		group.POST("updateCasbin", r.response.Handler()(system.Casbin.Update))
	}
	return r
}

func (r *casbin) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *casbin) PrivateWithoutRecord() interfaces.Router {
	group := r.router.Group("/casbin")
	{
		group.POST("getPolicyPathByAuthorityId", r.response.Handler()(system.Casbin.GetPolicyPathByAuthorityId))
	}
	return r
}
