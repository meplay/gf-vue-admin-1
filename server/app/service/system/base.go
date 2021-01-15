package service

import (
	"gf-vue-admin/app/api/response"
	"github.com/gogf/gf/frame/g"
	"github.com/mojocn/base64Captcha"
)

var Base = new(base)

type base struct {
	store base64Captcha.Store
}

func (b *base) Captcha() (result *response.Captcha, err error) {
	result = (*response.Captcha)(nil)
	b.store = base64Captcha.DefaultMemStore
	var imgHeight = g.Cfg("captcha").GetInt("captcha.ImageHeight")
	var imgWidth = g.Cfg("captcha").GetInt("captcha.ImageWidth")
	var keyLong = g.Cfg("captcha").GetInt("captcha.KeyLong")
	var driver = base64Captcha.NewDriverDigit(imgHeight, imgWidth, keyLong, 0.7, 80) // 字符,公式,验证码配置, 生成默认数字的driver
	cp := base64Captcha.NewCaptcha(driver, b.store)
	result.Id, result.Path, err = cp.Generate()
	return result, err
}
