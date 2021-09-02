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

func (a *_api) Init() {
	group := a.router.Group("/api")
	{
		group.POST("createApi", a.response.Handler()(system.Api.Create)) // 创建Api
		//group.POST("getApiById", a.response.Handler()(api.Api.First))          // 获取单条Api消息
		//group.POST("updateApi", a.response.Handler()(api.Api.Update))          // 更新api
		//group.POST("deleteApi", a.response.Handler()(api.Api.Delete))          // 删除Api
		//group.DELETE("deleteApisByIds", a.response.Handler()(api.Api.Deletes)) // 删除选中Api
		//group.POST("getApiList", a.response.Handler()(api.Api.GetList))        // 获取Api列表
		//group.POST("getAllApis", a.response.Handler()(api.Api.GetAllApis))     // 获取所有api
	}
}
