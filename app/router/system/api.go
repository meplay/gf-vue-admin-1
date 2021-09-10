package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type _api struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewApiRouter(router *ghttp.RouterGroup) *_api {
	return &_api{router: router, response: &response.Handler{}}
}

func (a *_api) Private() {
	group := a.router.Group("/api")
	{
		group.POST("createApi", a.response.Handler()(system.Api.Create))          // 创建api
		group.POST("getApiById", a.response.Handler()(system.Api.First))          // 根据id获取api
		group.POST("updateApi", a.response.Handler()(system.Api.Update))          // 更新api
		group.POST("deleteApi", a.response.Handler()(system.Api.Delete))          // 删除api
		group.DELETE("deleteApisByIds", a.response.Handler()(system.Api.Deletes)) // 批量删除api
		group.POST("getApiList", a.response.Handler()(system.Api.GetList))        // 分页获取api列表
		group.POST("getAllApis", a.response.Handler()(system.Api.GetAllApis))     // 获取所有api
	}
}
