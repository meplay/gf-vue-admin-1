package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type DictionaryDetailRouter struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewDictionaryDetailRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &DictionaryDetailRouter{router: router, response: &response.Handler{}}
}

func (d *DictionaryDetailRouter) Init() {
	var detail = d.router.Group("/sysDictionaryDetail")
	{
		detail.POST("createSysDictionaryDetail", d.response.Handler()(api.DictionaryDetail.Create))   // 新建DictionaryDetail
		detail.GET("findSysDictionaryDetail", d.response.Handler()(api.DictionaryDetail.First))       // 根据ID获取DictionaryDetail
		detail.PUT("updateSysDictionaryDetail", d.response.Handler()(api.DictionaryDetail.Update))    // 更新DictionaryDetail
		detail.DELETE("deleteSysDictionaryDetail", d.response.Handler()(api.DictionaryDetail.Delete)) // 删除DictionaryDetail
		detail.GET("getSysDictionaryDetailList", d.response.Handler()(api.DictionaryDetail.GetList))  // 获取DictionaryDetail列表
	}
}
