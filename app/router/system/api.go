package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/app/router/middleware"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*_api)(nil)

type _api struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewApiRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &_api{router: router, response: &response.Handler{}}
}

func (r *_api) Public() interfaces.Router {
	return r
}

func (r *_api) Private() interfaces.Router {
	group := r.router.Group("/api").Middleware(middleware.OperationRecord)
	{
		group.POST("createApi", r.response.Handler()(system.Api.Create))          // 创建api
		group.POST("updateApi", r.response.Handler()(system.Api.Update))          // 更新api
		group.POST("deleteApi", r.response.Handler()(system.Api.Delete))          // 删除api
		group.DELETE("deleteApisByIds", r.response.Handler()(system.Api.Deletes)) // 批量删除api
	}
	return r
}

func (r *_api) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *_api) PrivateWithoutRecord() interfaces.Router {
	group := r.router.Group("/api")
	{
		group.POST("getApiById", r.response.Handler()(system.Api.First))      // 根据id获取api
		group.POST("getApiList", r.response.Handler()(system.Api.GetList))    // 分页获取api列表
		group.POST("getAllApis", r.response.Handler()(system.Api.GetAllApis)) // 获取所有api
	}
	return r
}