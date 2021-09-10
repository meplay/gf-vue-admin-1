package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type dictionary struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewDictionaryRouter(router *ghttp.RouterGroup) *dictionary {
	return &dictionary{router: router, response: &response.Handler{}}
}

func (d *dictionary) Private() {
	group := d.router.Group("/sysDictionary")
	{
		group.POST("createSysDictionary", d.response.Handler()(system.Dictionary.Create))   // 创建系统字典
		group.GET("findSysDictionary", d.response.Handler()(system.Dictionary.First))       // 用id查询系统字典
		group.PUT("updateSysDictionary", d.response.Handler()(system.Dictionary.Update))    // 更新系统字典
		group.DELETE("deleteSysDictionary", d.response.Handler()(system.Dictionary.Delete)) // 删除系统字典
		group.GET("getSysDictionaryList", d.response.Handler()(system.Dictionary.GetList))  // 分页获取系统字典列表
	}
}
