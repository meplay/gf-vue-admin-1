package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/app/router/middleware"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*jwtBlacklist)(nil)

type jwtBlacklist struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewJwtBlacklistRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &jwtBlacklist{router: router, response: &response.Handler{}}
}

func (r *jwtBlacklist) Public() interfaces.Router {
	return r
}

func (r *jwtBlacklist) Private() interfaces.Router {
	group := r.router.Group("/jwt").Middleware(middleware.OperationRecord)
	{
		group.POST("jsonInBlacklist", r.response.Handler()(system.JwtBlacklist.JsonInBlacklist)) // jwt加入黑名单
	}
	return r
}

func (r *jwtBlacklist) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *jwtBlacklist) PrivateWithoutRecord() interfaces.Router {
	return r
}
