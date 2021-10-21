package interfaces

import "github.com/gogf/gf/net/ghttp"

// Plugin 插件模式接口化
type Plugin interface {
	// RouterPath 用户返回注册路由
	RouterPath() string
	// Register 注册路由
	Register(group *ghttp.RouterGroup)
}
