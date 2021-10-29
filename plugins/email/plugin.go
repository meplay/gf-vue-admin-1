package email

import (
	"github.com/gogf/gf/net/ghttp"
)

var Plugin = new(plugin)

type plugin struct{}

func (p *plugin) RouterPath() string {
	return "email"
}

func (p *plugin) PublicRegister(public *ghttp.RouterGroup) {
	NewEmailRouter(public).Public().PublicWithoutRecord()
}

func (p *plugin) PrivateRegister(private *ghttp.RouterGroup) {
	NewEmailRouter(private).Private().PrivateWithoutRecord()
}
