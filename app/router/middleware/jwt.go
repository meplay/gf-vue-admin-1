package middleware

import (
	"strconv"
	"time"

	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/auth"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go.uber.org/zap"
)

// JwtAuth jwt 中间件
// Author [SliverHorn](https://github.com/SliverHorn)
func JwtAuth(r *ghttp.Request) {
	// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
	token := r.Request.Header.Get("x-token")
	if token == "" {
		_ = r.Response.WriteJson(&response.Response{Code: 7, Data: g.Map{"reload": true}, Message: "未登录或非法访问!"})
		r.ExitAll()
		return
	}
	if system.JwtBlacklist.IsBlacklist(token) {
		_ = r.Response.WriteJson(&response.Response{Code: 7, Data: g.Map{"reload": true}, Message: "您的帐户异地登陆或令牌失效!"})
		r.ExitAll()
		return
	}
	_jwt := auth.NewJWT()
	claims, err := _jwt.ParseToken(token)
	if err != nil {
		if err == auth.TokenExpired {
			_ = r.Response.WriteJson(&response.Response{Error: err, Data: g.Map{"reload": true}, Message: "授权已过期!"})
			r.ExitAll()
			return
		}
		_ = r.Response.WriteJson(&response.Response{Error: err, Data: g.Map{"reload": true}, Message: "解析token失败!"})
		r.ExitAll()
		return
	}

	//if _, err = system.User.Find(&request.UserFind{Uuid: claims.Uuid}); err != nil {
	//	_ = system.JwtBlacklist.JwtToBlacklist(token)
	//	_ = r.Response.WriteJson(&response.Response{Error: err, Data: g.Map{"reload": true}, Message: "用户不存在!"})
	//	r.ExitAll()
	//} // 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

	if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
		claims.ExpiresAt = time.Now().Unix() + global.Config.Jwt.ExpiresTime
		newToken, _ := _jwt.CreateTokenByOldToken(token, claims)
		newClaims, _ := _jwt.ParseToken(newToken)
		r.Response.Header().Set("new-token", newToken)
		r.Response.Header().Set("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		if global.Config.System.UseMultipoint {
			_token, _err := system.JwtBlacklist.GetRedisJWT(newClaims.Username)
			if _err != nil {
				zap.L().Error("get redis jwt failed", zap.Error(_err))
			} else { // 当之前的取成功时才进行拉黑操作
				_ = system.JwtBlacklist.JwtToBlacklist(_token)
			}
			_ = system.JwtBlacklist.SetRedisJWT(newToken, newClaims.Username) // 无论如何都要记录当前的活跃状态
		}
	}
	r.SetCtxVar("claims", claims)
	r.Middleware.Next()
}
