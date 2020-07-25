package boot

import (
	"server/app/middleware"
	"server/library/global"
	"server/router"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

func InitializeRunServer() {
	global.GFVA_SERVER = g.Server()
	global.GFVA_SERVER.Use(middleware.Error)
	global.GFVA_SERVER.SetReadTimeout(10 * time.Second)
	global.GFVA_SERVER.SetWriteTimeout(10 * time.Second)
	global.GFVA_SERVER.SetMaxHeaderBytes(1 << 20)
	router.InitializeRouters()
	global.GFVA_SERVER.Plugin(&swagger.Swagger{})
	global.GFVA_SERVER.Run()
}
