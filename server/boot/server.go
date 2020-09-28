package boot

import (
	"server/app/middleware"
	"server/router"
	"time"

	"github.com/gogf/gf/frame/g"
)

func InitializeRunServer() {
	s := g.Server()
	s.Use(middleware.Error)
	s.SetReadTimeout(10 * time.Second)
	s.SetWriteTimeout(10 * time.Second)
	s.SetMaxHeaderBytes(1 << 20)
	s.SetIndexFolder(true)
	if g.Cfg().GetString("system.OssType") == "local" {
		s.AddStaticPath("/"+g.Cfg("oss").GetString("local.LocalPath"), g.Cfg("oss").GetString("local.LocalPath"))
	}
	s.AddStaticPath("/form-generator", "public/page")
	router.InitializeRouters()
	s.Run()
}
