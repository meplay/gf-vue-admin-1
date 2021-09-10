package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/response"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/mojocn/base64Captcha"
)

var (
	Store   base64Captcha.Store
	Captcha = new(captcha)
)

type captcha struct{}

// Captcha 验证码生成
// 字符,公式,验证码配置
// 生成默认数字的driver
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *captcha) Captcha() (*response.Captcha, error) {
	driver := base64Captcha.NewDriverDigit(global.Config.Captcha.ImgHeight, global.Config.Captcha.ImgWidth, global.Config.Captcha.KeyLong, 0.7, 80)
	Store = base64Captcha.DefaultMemStore
	// Store = store.NewRedisStore() // redis 缓存 base64Captcha库数据
	cp := base64Captcha.NewCaptcha(driver, Store)
	id, b64s, err := cp.Generate()
	return &response.Captcha{PicPath: b64s, CaptchaId: id}, err
}
