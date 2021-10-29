package boot

import (
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/plugins/email"
	"github.com/gogf/gf/net/ghttp"
)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) PublicInitialize(public *ghttp.RouterGroup) {
	interfaces.PublicInit(public, email.Plugin)
}

func (p *plugin) PrivateInitialize(private *ghttp.RouterGroup) {
	interfaces.PrivateInit(private, email.Plugin)
}
