package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/app/router/middleware"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*dictionary)(nil)

type dictionary struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewDictionaryRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &dictionary{router: router, response: &response.Handler{}}
}

func (r *dictionary) Public() interfaces.Router {
	return r
}

func (r *dictionary) Private() interfaces.Router {
	group := r.router.Group("/sysDictionary").Middleware(middleware.OperationRecord)
	{
		group.POST("createSysDictionary", r.response.Handler()(system.Dictionary.Create))   // 创建系统字典
		group.PUT("updateSysDictionary", r.response.Handler()(system.Dictionary.Update))    // 更新系统字典
		group.DELETE("deleteSysDictionary", r.response.Handler()(system.Dictionary.Delete)) // 删除系统字典
	}
	return r
}

func (r *dictionary) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *dictionary) PrivateWithoutRecord() interfaces.Router {
	group := r.router.Group("/sysDictionary")
	{
		group.GET("findSysDictionary", r.response.Handler()(system.Dictionary.First))      // 用id查询系统字典
		group.GET("getSysDictionaryList", r.response.Handler()(system.Dictionary.GetList)) // 分页获取系统字典列表
	}
	return r
}
