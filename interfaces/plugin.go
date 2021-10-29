package interfaces

import "github.com/gogf/gf/net/ghttp"

// Plugin 插件模式接口化
type Plugin interface {
	// RouterPath 用户返回注册路由
	RouterPath() string
	// Register 注册路由
	Register(group *ghttp.RouterGroup)
}

// PluginInit 注册
func PluginInit(group *ghttp.RouterGroup, plugins ...Plugin) {
	for i := 0; i < len(plugins); i++ {
		plugin := group.Group(plugins[i].RouterPath())
		plugins[i].Register(plugin)
	}
}
