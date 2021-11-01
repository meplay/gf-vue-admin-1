package interfaces

import (
	"github.com/gogf/gf/net/ghttp"
	"strings"
)

// Plugin 插件模式接口化
type Plugin interface {
	// RouterPath 用户返回注册路由
	RouterPath() string
	// PublicRouterGroup 公开路由组
	PublicRouterGroup(group *ghttp.RouterGroup) PublicRouter
	// PrivateRouterGroup 私有路由组注册
	PrivateRouterGroup(group *ghttp.RouterGroup) PrivateRouter
}

// PublicInit 公开路由注册初始化
func PublicInit(group *ghttp.RouterGroup, plugins ...Plugin) {
	for i := 0; i < len(plugins); i++ {
		path := plugins[i].RouterPath()
		if !strings.Contains(path, "/") {
			path = "/" + path
		}
		plugin := group.Group(path)
		if router := plugins[i].PublicRouterGroup(plugin); router != nil {
			router.Public().PublicWithoutRecord()
		}
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
		if router := plugins[i].PrivateRouterGroup(plugin); router != nil {
			router.Private().PrivateWithoutRecord()
		}
	}
}
