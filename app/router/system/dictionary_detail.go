package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/app/router/middleware"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*dictionaryDetail)(nil)

type dictionaryDetail struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewDictionaryDetailRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &dictionaryDetail{router: router, response: &response.Handler{}}
}

func (r *dictionaryDetail) Public() interfaces.Router {
	return r
}

func (r *dictionaryDetail) Private() interfaces.Router {
	group := r.router.Group("/sysDictionaryDetail").Middleware(middleware.OperationRecord)
	{
		group.POST("createSysDictionaryDetail", r.response.Handler()(system.DictionaryDetail.Create))   // 创建系统字典详情
		group.PUT("updateSysDictionaryDetail", r.response.Handler()(system.DictionaryDetail.Update))    // 更新系统详情字典
		group.DELETE("deleteSysDictionaryDetail", r.response.Handler()(system.DictionaryDetail.Delete)) // 删除系统详情字典
	}
	return r
}

func (r *dictionaryDetail) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *dictionaryDetail) PrivateWithoutRecord() interfaces.Router {
	group := r.router.Group("/sysDictionaryDetail")
	{
		group.GET("findSysDictionaryDetail", r.response.Handler()(system.DictionaryDetail.First))      // 用id查询系统详情字典
		group.GET("getSysDictionaryDetailList", r.response.Handler()(system.DictionaryDetail.GetList)) // 分页获取系统字典详情列表
	}
	return r
}
