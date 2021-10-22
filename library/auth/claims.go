package auth

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/gogf/gf/net/ghttp"
	"go.uber.org/zap"
)

var Claims = new(claims)

type claims struct{}

func (c *claims) GetUserInfo(r *ghttp.Request) *request.CustomClaims {
	data := r.GetCtxVar("claims")
	var _claims request.CustomClaims
	if err := data.Struct(&_claims); err != nil {
		zap.L().Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!", zap.Error(err))
	}
	return &_claims
}
