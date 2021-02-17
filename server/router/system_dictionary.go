package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type DictionaryRouter struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewDictionaryRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &DictionaryRouter{router: router, response: &response.Handler{}}
}

func (d *DictionaryRouter) Init() {
	var dictionary = d.router.Group("/sysDictionary")
	{
		dictionary.POST("createSysDictionary", d.response.Handler()(api.Dictionary.Create))   // 新建Dictionary
		dictionary.GET("findSysDictionary", d.response.Handler()(api.Dictionary.First))       // 根据ID获取Dictionary
		dictionary.PUT("updateSysDictionary", d.response.Handler()(api.Dictionary.Update))    // 更新Dictionary
		dictionary.DELETE("deleteSysDictionary", d.response.Handler()(api.Dictionary.Delete)) // 删除Dictionary
		dictionary.GET("getSysDictionaryList", d.response.Handler()(api.Dictionary.GetList))  // 获取Dictionary列表
	}
}
