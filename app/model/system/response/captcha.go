package response

type Captcha struct {
	PicPath   string `json:"picPath"`
	CaptchaId string `json:"captchaId"`
}
