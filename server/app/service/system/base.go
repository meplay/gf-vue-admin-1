package service

import (
	"database/sql"
	"gf-vue-admin/app/model/system/response"
	"gf-vue-admin/library"
	"gf-vue-admin/library/global"
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

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 生成二维码的信息
func (b *base) Captcha() (result *response.Captcha, err error) {
	var info response.Captcha
	var driver = base64Captcha.NewDriverDigit(global.Config.Captcha.ImageHeight, global.Config.Captcha.ImageWidth, global.Config.Captcha.KeyLong, 0.7, 80) // 字符,公式,验证码配置, 生成默认数字的driver
	var captcha = base64Captcha.NewCaptcha(driver, Store)
	info.Id, info.Path, err = captcha.Generate()
	return &info, err
}