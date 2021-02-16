package api

import (
	"gf-vue-admin/app/api/response"
	service "gf-vue-admin/app/service/system"
	"github.com/gogf/gf/net/ghttp"
)

var Base = new(base)

type base struct{}

// @Tags SystemBase
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func (b *base) Captcha(r *ghttp.Request) *response.Response {
	if result, err := service.Base.Captcha(); err != nil {
		return &response.Response{Data: result, MessageCode: response.ErrorCaptcha}
	} else {
		return &response.Response{Data: result, MessageCode: response.SuccessCaptcha}
	}
}
