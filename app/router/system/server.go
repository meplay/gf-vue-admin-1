package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/app/router/middleware"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*server)(nil)

type server struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewServerRouter(router *ghttp.RouterGroup) *server {
	return &server{router: router, response: &response.Handler{}}
}

func (r *server) Public() interfaces.Router {
	return r
}

func (r *server) Private() interfaces.Router {
	group := r.router.Group("/system").Middleware(middleware.OperationRecord)
	{
		group.POST("reloadSystem", r.response.Handler()(system.Server.ReloadSystem))       // 重启服务
		group.POST("setSystemConfig", r.response.Handler()(system.Server.SetSystemConfig)) // 设置配置文件内容
	}
	return r
}

func (r *server) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *server) PrivateWithoutRecord() interfaces.Router {
	group := r.router.Group("/system")
	{
		group.POST("getServerInfo", r.response.Handler()(system.Server.GetServerInfo))     // 获取服务器信息
		group.POST("getSystemConfig", r.response.Handler()(system.Server.GetSystemConfig)) // 获取配置文件内容
	}
	return r
}
