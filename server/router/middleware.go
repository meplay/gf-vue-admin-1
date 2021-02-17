package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	service "gf-vue-admin/app/service/system"
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 允许接口跨域请求
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 处理panic产生的错误
func Error(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		g.Log("exception").Error(err) // 记录到自定义错误日志文件
		r.Response.ClearBuffer()      //返回固定的友好信息
		//global.FailWithMessage(r, "服务器居然开小差了，请稍后再试吧！")
	}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 验证token有效性
func JwtAuth(r *ghttp.Request) {
	api.GfJWTMiddleware.MiddlewareFunc()(r)
	_jwt, err := api.GfJWTMiddleware.ParseToken(r) // 解析token
	if err != nil {
		if err == jwt.ErrExpiredToken {
			_ = r.Response.WriteJson(&response.Response{Code: 7, Data: g.Map{"reload": true}, Message: "授权已过期!"})
			r.ExitAll()
		}
		_ = r.Response.WriteJson(&response.Response{Code: 7, Data: g.Map{"reload": true}, Error: err})
		r.ExitAll()
	}
	if _jwt != nil {
		token := _jwt.Raw
		if service.JwtBlacklist.IsBlacklist(token) {
			_ = r.Response.WriteJson(&response.Response{Code: 7, Data: g.Map{"reload": true}, Message: "您的帐户异地登陆或令牌失效!"})
			r.ExitAll()
		}
		var claims = gconv.Map(_jwt.Claims)
		r.SetParam("claims", _jwt.Claims)
		r.SetParam("admin_authority_id", claims["admin_authority_id"])
		if g.Cfg("system").GetBool("system.UseMultipoint") {
			if !service.JwtBlacklist.ValidatorRedisToken(gconv.String(claims["admin_uuid"]), token) {
				_ = r.Response.WriteJson(&response.Response{Code: 7, Data: g.Map{"reload": true}, Message: "Token鉴权失败!"})
				r.Exit()
			}
		}
	}

	r.Middleware.Next()
}
