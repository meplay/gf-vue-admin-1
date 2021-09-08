package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	_response "github.com/flipped-aurora/gf-vue-admin/app/model/system/response"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
	"github.com/mojocn/base64Captcha"
)

var (
	store = base64Captcha.DefaultMemStore
	Base  = new(base)
)

type base struct{}

// Captcha
// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":""}"
// @Success 200 {object} response.Response{data=response.Captcha} "验证码获取成功!"
// @Router /base/captcha [post]
func (b *base) Captcha(r *ghttp.Request) *response.Response {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.Config.Captcha.ImgHeight, global.Config.Captcha.ImgWidth, global.Config.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		return &response.Response{Error: err, Message: "验证码获取失败!"}
	}
	return &response.Response{Data: _response.Captcha{PicPath: b64s, CaptchaId: id}, Message: "验证码获取成功!"}
}

// Login
// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.UserLogin true "请求参数"
// @Success 200 {object} response.Response{data=response.UserLogin} "登录成功!"
// @Router /base/login [post]
func (b *base) Login(r *ghttp.Request) *response.Response {
	var info request.UserLogin
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if store.Verify(info.CaptchaId, info.Captcha, true) {
		data, err := system.User.Login(&info)
		if err != nil {
			return &response.Response{Error: err, Message: "登录失败!"}
		}
		return &response.Response{Data: data, Message: "登录成功!"}
	} else {
		return &response.Response{Code: 7, Message: "验证码错误!"}
	}
}
