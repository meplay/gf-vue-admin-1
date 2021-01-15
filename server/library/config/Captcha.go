package config

type Captcha struct {
	KeyLong     int `json:"key_long"`
	ImageWidth  int `json:"image_width"`
	ImageHeight int `json:"image_height"`
}
