package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// AdminRegister Administrator account registration
// AdminRegister 管理员注册账号
func AdminRegister(r *ghttp.Request) {
	var R *request.AdminRegister
	if err := r.Parse(&R); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.AdminRegister(R); err != nil {
		global.FailWithMessage(r, err.Error())
		r.ExitAll()
	}
	global.OkDetailed(r, g.Map{}, "注册成功!")
}

// Captcha Generate verification code
// Captcha 生成验证码
func Captcha(r *ghttp.Request) {
	id, b64s, err := service.Captcha()
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，err:%v", err))
		r.Exit()
	}
	global.OkDetailed(r, g.Map{"captchaId": id, "picPath": b64s}, "验证码获取成功")
}
