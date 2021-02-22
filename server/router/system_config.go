package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type ConfigRouter struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewConfigRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &ConfigRouter{router: router, response: &response.Handler{}}
}

func (c *ConfigRouter) Init() {
	var config = c.router.Group("/system")
	{
		config.POST("getSystemConfig", c.response.Handler()(api.Config.GetConfig)) // 获取配置文件内容
		config.POST("setSystemConfig", c.response.Handler()(api.Config.SetConfig)) // 设置配置文件内容
		//config.POST("getServerInfo", c.response.Handler()(api.Config.GetServerInfo)) // 获取服务器信息
	}
}
