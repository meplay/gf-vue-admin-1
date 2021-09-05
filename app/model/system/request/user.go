package request

type Login struct {
	Captcha   string `json:"captcha"`
	Username  string `p:"username"`
	Password  string `p:"password"`
	CaptchaId string `json:"captchaId"`
}
