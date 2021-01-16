package service

import (
	"gf-vue-admin/app/api/response"
	"gf-vue-admin/library/global"
	"github.com/mojocn/base64Captcha"
)

var Base = new(base)
var Store = base64Captcha.DefaultMemStore

type base struct{}

func (b *base) Login() {
	//var entity = (*model.Admin)(nil)

}

func (b *base) Captcha() (result *response.Captcha, err error) {
	result = (*response.Captcha)(nil)
	var driver = base64Captcha.NewDriverDigit(global.Config.Captcha.ImageHeight, global.Config.Captcha.ImageWidth, global.Config.Captcha.KeyLong, 0.7, 80) // 字符,公式,验证码配置, 生成默认数字的driver
	cp := base64Captcha.NewCaptcha(driver, Store)
	result.Id, result.Path, err = cp.Generate()
	return result, err
}
