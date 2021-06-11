package router

import (
	"flipped-aurora/gf-vue-admin/server/interfaces"
	extra "flipped-aurora/gf-vue-admin/server/router/extra"
	"flipped-aurora/gf-vue-admin/server/router/internal"
	system "flipped-aurora/gf-vue-admin/server/router/system"
	workflow "flipped-aurora/gf-vue-admin/server/router/workflow"
	"github.com/gogf/gf/frame/g"
)

var Routers = new(routers)

type routers struct{}

func (r *routers) Init() {
	public := g.Server().Group("")
	interfaces.RouterInit( // 无需鉴权中间件
		system.NewBaseGroup(public),
	)

	private := g.Server().Group("").Middleware(internal.Middleware.JwtAuth, internal.Middleware.CasbinRbac)
	interfaces.RouterInit( // 需要Jwt鉴权, casbin鉴权
		system.NewApiRouter(private),
		system.NewMenuRouter(private),
		system.NewEmailRouter(private),
		system.NewAdminRouter(private),
		system.NewSystemRouter(private),
		system.NewCasbinRouter(private),
		system.NewGenerateRouter(private),
		system.NewAuthorityRouter(private),
		system.NewDictionaryRouter(private),
		system.NewJwtBlacklistRouter(private),
		system.NewOperationRecordRouter(private),
		system.NewDictionaryDetailRouter(private),

		extra.NewFileRouter(private),
		extra.NewSimpleUploaderRouter(private),

		workflow.NewWorkflowRouter(private),
	)
}
