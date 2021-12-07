package boot

import (
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
	"go.uber.org/zap"
)

var Server = new(_server)

type _server struct{}

func (s *_server) Initialize() {
	server := g.Server()
	address := g.Cfg().GetString("server.address")
	server.SetIndexFolder(true)
	server.AddStaticPath("/form-generator", "public/page")
	Routers.PublicRouter().PrivateRouter()
	server.SetPort()
	server.Plugin(&swagger.Swagger{})
	zap.L().Info(fmt.Sprintf(`
	欢迎使用 Gf-Vue-Admin
	当前版本:V1.0.1
	加群方式:微信号：SliverHorn QQ群：1040044540
	默认自动化文档地址:http://127.0.0.1%s/swagger
	默认前端文件运行地址:http://127.0.0.1:8080
	如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gf-vue-admin.com/docs/coffee
`, address))
	server.Run()
}
