package boot

import (
	"server/app/middleware"
	"server/router"
	"time"

	"github.com/gogf/gf/frame/g"
)

func InitializeRunServer() {
	g.Server() = g.Server()
	g.Server().Use(middleware.Error)
	g.Server().SetReadTimeout(10 * time.Second)
	g.Server().SetWriteTimeout(10 * time.Second)
	g.Server().SetMaxHeaderBytes(1 << 20)
	router.InitializeRouters()
	g.Server().Run()
}
