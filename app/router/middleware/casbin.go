package middleware

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// Casbin Casbin拦截器
func Casbin(r *ghttp.Request) {
	claims := r.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	url := r.Request.URL.RequestURI()  // 获取请求的URI
	method := r.Request.Method         // 获取请求方法
	authorityId := waitUse.AuthorityId // 获取用户的角色
	e := system.Casbin.Casbin()
	success, _ := e.Enforce(authorityId, url, method)
	if global.Config.System.Env == "develop" || success {
		r.Middleware.Next()
	} else {
		_ = r.Response.WriteJson(&response.Response{Code: 7, Message: "权限不足!"})
		r.ExitAll()
		return
	} // 判断策略中是否存在
	return
}
