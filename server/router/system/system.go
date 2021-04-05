package router

import (
	"gf-vue-admin/library/response"
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"gf-vue-admin/router/internal"
	"github.com/gogf/gf/net/ghttp"
)

type system struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewSystemRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &system{router: router, response: &response.Handler{}}
}

func (c *system) Init() {
	group := c.router.Group("/system").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("getSystemConfig", c.response.Handler()(api.System.GetConfig))   // 获取配置文件内容
		group.POST("setSystemConfig", c.response.Handler()(api.System.SetConfig))   // 设置配置文件内容
		group.POST("getServerInfo", c.response.Handler()(api.System.GetServerInfo)) // 获取服务器信息
		group.POST("reloadSystem", c.response.Handler()(api.System.ReloadSystem))   // 重启服务
	}
}
