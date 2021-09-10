package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Captcha = new(captcha)

type captcha struct{}

// Captcha
// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":""}"
// @Success 200 {object} response.Response{data=response.Captcha} "验证码获取成功!"
// @Router /base/captcha [post]
func (a *captcha) Captcha(r *ghttp.Request) *response.Response {
	data, err := system.Captcha.Captcha()
	if err != nil {
		return &response.Response{Error: err, Message: "验证码获取失败!"}
	}
	return &response.Response{Data: data, Message: "验证码获取成功!"}
}
