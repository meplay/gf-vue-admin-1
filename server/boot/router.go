package boot

import (
	"gf-vue-admin/router"
	"github.com/gogf/gf/net/ghttp"
)

var Routers = new(routers)

type routers struct{}

func (r *routers) Init(server *ghttp.Server) {
	//var public = server.Group("")
	//{
	//
	//}
	//var private = g.Server().Group("").Middleware(middleware.JwtAuth, middleware.CasbinMiddleware)
	var private = server.Group("")
	{ // 需要Jwt鉴权, casbin鉴权
		router.NewAdminRouter(private)
	}
}
