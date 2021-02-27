package router

import (
	"github.com/gogf/gf/frame/g"
)

var Routers = new(routers)

type routers struct{}

func (r *routers) Init() {
	public := g.Server().Group("")
	{
		NewBaseGroup(public).Init()
	}
	private := g.Server().Group("").Middleware(JwtAuth, CasbinRbac)
	{ // 需要Jwt鉴权, casbin鉴权
		NewApiRouter(private).Init()
		NewAdminRouter(private).Init()
		NewMenuRouter(private).Init()
		NewConfigRouter(private).Init()
		NewCasbinRouter(private).Init()
		NewGenerateRouter(private).Init()
		NewAuthorityRouter(private).Init()
		NewDictionaryRouter(private).Init()
		NewJwtBlacklistRouter(private).Init()
		NewOperationRecordRouter(private).Init()
		NewDictionaryDetailRouter(private).Init()

		NewFileRouter(private).Init()
	}
}
