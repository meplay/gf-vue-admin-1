package boot

import (
	"fmt"
	"gf-vue-admin/library/global"
	"gf-vue-admin/library/utils"
	"gf-vue-admin/router"
	"github.com/gogf/swagger"
	"time"

	"github.com/gogf/gf/frame/g"
)

var Server = new(_server)

type _server struct{}

func (s *_server) Initialize() {
	var server = g.Server()
	server.SetReadTimeout(10 * time.Second)
	server.SetWriteTimeout(10 * time.Second)
	server.SetMaxHeaderBytes(1 << 20)
	server.SetIndexFolder(true)
	if global.Config.System.OssType == "local" {
		_ = utils.CreateDir(g.Cfg("oss").GetString("local.LocalPath"))
		server.AddStaticPath("/"+g.Cfg("oss").GetString("local.LocalPath"), g.Cfg("oss").GetString("local.LocalPath"))
	}
	server.AddStaticPath("/form-generator", "public/page")
	router.Routers.Init()
	fmt.Printf(`
	欢迎使用 GoFrame-Vue-Admin
	当前版本:V2.3.9
	加群方式:微信号：SliverHorn QQ群：1040044540
	默认自动化文档地址:http://127.0.0.1%s/swagger
	默认前端文件运行地址:http://127.0.0.1:8080
	如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/docs/coffee
	
`, g.Cfg().GetString("server.Address"))
	server.Plugin(&swagger.Swagger{})
	server.SetPort()
	server.Run()
}
