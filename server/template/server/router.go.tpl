package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type {{.StructName}}Router struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func New{{.StructName}}Router(router *ghttp.RouterGroup) interfaces.Router {
	return &{{.StructName}}Router{router: router, response: &response.Handler{}}
}

func (r *{{.StructName}}Router) Init() {
	{{.Abbreviation}} := r.router.Group("/{{.Abbreviation}}")
	{
		{{.Abbreviation}}.POST("create", r.response.Handler()(api.{{.StructName}}.Create))             // 新建{{.StructName}}
		{{.Abbreviation}}.GET("first", r.response.Handler()(api.{{.StructName}}.First))                // 根据ID获取{{.StructName}}
		{{.Abbreviation}}.PUT("update", r.response.Handler()(api.{{.StructName}}.Update))              // 更新{{.StructName}}
		{{.Abbreviation}}.DELETE("delete", r.response.Handler()(api.{{.StructName}}.Delete))           // 删除{{.StructName}}
		{{.Abbreviation}}.DELETE("deletes", r.response.Handler()(api.{{.StructName}}.Deletes))         // 批量删除{{.StructName}}
		{{.Abbreviation}}.POST("getList", r.response.Handler()(api.{{.StructName}}.GetList))           // 获取{{.StructName}}列表
	}
}