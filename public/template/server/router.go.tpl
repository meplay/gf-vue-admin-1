package example

import (
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*{{.Abbreviation}})(nil)

type {{.Abbreviation}} struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewUserRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &{{.Abbreviation}}{router: router, response: &response.Handler{}}
}

func (r *{{.Abbreviation}}) Public() interfaces.Router {
	return r
}

func (r *{{.Abbreviation}}) Private() interfaces.Router {
	group := r.router.Group("/api").Middleware(middleware.OperationRecord)
	{
		group.POST("create", r.response.Handler()(example.{{.StructName}}.Create))     // 创建{{.Description}}
		group.POST("update", r.response.Handler()(example.{{.StructName}}.Update))     // 更新{.Description}}
		group.POST("delete", r.response.Handler()(example.{{.StructName}}.Delete))     // 删除{.Description}}
		group.DELETE("deletes", r.response.Handler()(example.{{.StructName}}.Deletes)) // 批量删除{.Description}}
	}
	return r
}

func (r *{{.Abbreviation}}) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *{{.Abbreviation}}) PrivateWithoutRecord() interfaces.Router {
	group := r.router.Group("/{{.Abbreviation}}")
	{
		group.GET("find", r.response.Handler()(example.{{.StructName}}.Find))       // 根据id获取{.Description}}
		group.POST("getList", r.response.Handler()(example.{{.StructName}}.GetList)) // 分页获取{.Description}}列表
	}
	return r
}