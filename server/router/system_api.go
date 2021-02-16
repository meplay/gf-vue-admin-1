package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type ApiRouter struct {
	router    *ghttp.RouterGroup
	response *response.Handler
}

func NewApiRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &ApiRouter{router: router, response: &response.Handler{}}
}

func (a *ApiRouter) Init() {
	var apis = a.router.Group("/api")
	{
		apis.POST("createApi", a.response.Handler()(api.Api.Create))      // 创建Api
		apis.POST("getApiById", a.response.Handler()(api.Api.First))      // 获取单条Api消息
		apis.POST("updateApi", a.response.Handler()(api.Api.Update))      // 更新api
		apis.POST("deleteApi", a.response.Handler()(api.Api.Delete))      // 删除Api
		apis.POST("getApiList", a.response.Handler()(api.Api.GetList))    // 获取Api列表
		apis.POST("getAllApis", a.response.Handler()(api.Api.GetAllApis)) // 获取所有api
	}
}