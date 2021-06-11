package service

import (
	"database/sql"
	"flipped-aurora/gf-vue-admin/server/app/model/system/response"
	"flipped-aurora/gf-vue-admin/server/library"
	"flipped-aurora/gf-vue-admin/server/library/global"
	"github.com/mojocn/base64Captcha"
	"gorm.io/gorm"
)

var (
	Store base64Captcha.Store
	Base  = new(base)
)

func init() {
	if global.Config.Captcha.CaptchaInRedis {
		Store = library.RedisStore
	} else {
		Store = base64Captcha.DefaultMemStore
	}
}

type base struct {
	db  *gorm.DB
	err error
	sql *sql.DB
}

// Captcha 生成二维码的信息
// Author: [SliverHorn](https://github.com/SliverHorn)
func (b *base) Captcha() (result *response.Captcha, err error) {
	var info response.Captcha
	var driver = base64Captcha.NewDriverDigit(global.Config.Captcha.ImageHeight, global.Config.Captcha.ImageWidth, global.Config.Captcha.KeyLong, 0.7, 80) // 字符,公式,验证码配置, 生成默认数字的driver
	var captcha = base64Captcha.NewCaptcha(driver, Store)
	info.Id, info.Path, err = captcha.Generate()
	return &info, err
}