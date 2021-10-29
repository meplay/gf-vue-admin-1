package boot

import (
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) PublicInitialize(public *ghttp.RouterGroup) {
	interfaces.PluginInit(public)
}

func (p *plugin) PrivateInitialize(private *ghttp.RouterGroup) {
	interfaces.PluginInit(private)
}
