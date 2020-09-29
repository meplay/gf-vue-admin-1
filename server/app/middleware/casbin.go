package middleware

import (
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

// CasbinMiddleware The interceptor
// CasbinMiddleware 拦截器
func CasbinMiddleware(r *ghttp.Request) {
	// 获取请求的URI
	obj := r.Request.URL.RequestURI()
	// 获取请求方法
	act := r.Request.Method
	// 获取用户的角色
	sub := r.GetParam("admin_authority_id")
	e := service.Casbin()
	// 判断策略中是否存在
	success, _ := e.Enforce(sub, obj, act)
	if g.Cfg("system").GetString("system.Env") == "develop" || success {
		r.Middleware.Next()
	} else {
		global.Result(r, global.ERROR, g.Map{}, "权限不足")
		r.ExitAll()
	}
}
