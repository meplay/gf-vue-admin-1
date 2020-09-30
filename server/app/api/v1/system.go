package v1

import (
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

// GetSystemConfig 获取配置文件内容
func GetSystemConfig(r *ghttp.Request) {
	config := service.GetSystemConfig()
	global.OkWithData(r, g.Map{"config": config})
}
