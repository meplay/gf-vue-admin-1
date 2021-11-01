package email

import (
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) RouterPath() string {
	return "email"
}

func (p *plugin) PublicRouterGroup(public *ghttp.RouterGroup) interfaces.PublicRouter {
	return nil
}

func (p *plugin) PrivateRouterGroup(private *ghttp.RouterGroup) interfaces.PrivateRouter {
	return NewEmailPrivateRouter(private)
}
