package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type dictionaryDetail struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewDictionaryDetailRouter(router *ghttp.RouterGroup) *dictionaryDetail {
	return &dictionaryDetail{router: router, response: &response.Handler{}}
}

func (d *dictionaryDetail) Init() {
	group := d.router.Group("/sysDictionaryDetail")
	{
		group.POST("createSysDictionaryDetail", d.response.Handler()(system.DictionaryDetail.Create))   // 创建系统字典详情
		group.GET("findSysDictionaryDetail", d.response.Handler()(system.DictionaryDetail.First))       // 用id查询系统详情字典
		group.PUT("updateSysDictionaryDetail", d.response.Handler()(system.DictionaryDetail.Update))    // 更新系统详情字典
		group.DELETE("deleteSysDictionaryDetail", d.response.Handler()(system.DictionaryDetail.Delete)) // 删除系统详情字典
		group.GET("getSysDictionaryDetailList", d.response.Handler()(system.DictionaryDetail.GetList))  // 分页获取系统字典详情列表
	}
}
