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
	//s.SetIndexFolder(true)
	s.SetServerRoot("/public/page")
	//s.AddStaticPath("/form-generator", "/public/page")
	s.SetRewrite("/form-generator", "/index.html")
	router.InitializeRouters()
	s.Run()
}
