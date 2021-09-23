package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*autoCode)(nil)

type autoCode struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewAutoCodeRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &autoCode{router: router, response: &response.Handler{}}
}

func (r *autoCode) Public() interfaces.Router {
	return r
}

func (r *autoCode) Private() interfaces.Router {
	return r
}

func (r *autoCode) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *autoCode) PrivateWithoutRecord() interfaces.Router {
	group := r.router.Group("/autoCode")
	{
		group.GET("getDB", r.response.Handler()(system.AutoCode.GetDbs))          // 获取数据库
		group.GET("getTables", r.response.Handler()(system.AutoCode.GetTables))   // 获取对应数据库的表
		group.GET("getColumn", r.response.Handler()(system.AutoCode.GetColumns))  // 获取指定表所有字段信息
		group.POST("preview", r.response.Handler()(system.AutoCode.Preview))      // 获取自动创建代码预览
		group.POST("createTemp", r.response.Handler()(system.AutoCode.Create))    // 创建自动化代码
	}
	return r
}
