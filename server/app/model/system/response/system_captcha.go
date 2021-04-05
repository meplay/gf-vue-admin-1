package response

type Captcha struct {
	Id   string `json:"captchaId"`
	Path string `json:"picPath"`
}
