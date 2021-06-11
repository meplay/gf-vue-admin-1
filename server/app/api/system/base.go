package api

import (
	service "flipped-aurora/gf-vue-admin/server/app/service/system"
	"flipped-aurora/gf-vue-admin/server/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Base = new(base)

type base struct{}
// Captcha
// Author: [SliverHorn](https://github.com/SliverHorn)
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