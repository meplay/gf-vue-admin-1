package interfaces

import (
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

// Plugin 插件模式接口化
type Plugin interface {
	// RouterPath 用户返回注册路由
	RouterPath() string
	// PublicRegister 公开路由组
	PublicRegister(group *ghttp.RouterGroup)
	// PrivateRegister 私有路由组注册
	PrivateRegister(group *ghttp.RouterGroup)
}

// PublicInit 公开路由注册初始化
func PublicInit(group *ghttp.RouterGroup, plugins ...Plugin) {
	for i := 0; i < len(plugins); i++ {
		path := plugins[i].RouterPath()
		if !strings.Contains(path, "/") {
			path = "/" + path
		}
		plugin := group.Group(path)
		plugins[i].PublicRegister(plugin)
	}
}

// PrivateInit 私有路由组注册初始化
func PrivateInit(group *ghttp.RouterGroup, plugins ...Plugin) {
	for i := 0; i < len(plugins); i++ {
		path := plugins[i].RouterPath()
		if !strings.Contains(path, "/") {
			path = "/" + path
		}
		plugin := group.Group(path)
		plugins[i].PrivateRegister(plugin)
	}
}
