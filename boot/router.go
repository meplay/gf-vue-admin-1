package boot

import (
	"github.com/flipped-aurora/gf-vue-admin/router/system"
	"github.com/gogf/gf/frame/g"
)

var Routers = new(_router)

type _router struct{}

// Initialize 路由初始化
// Author [SliverHorn](https://github.com/SliverHorn)
func (r *_router) Initialize() {
	public := g.Server().Group("")
	{
		system.NewApiRouter(public)
	} // 无需鉴权中间件
	private := g.Server().Group("")
	{
		system.NewApiRouter(private)
	} // 需要Jwt鉴权, casbin鉴权
}
