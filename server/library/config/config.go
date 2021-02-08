package config

type Config struct {
	Jwt            Jwt            `json:"jwt"`
	Oss            Oss            `json:"oss"`
	Email          Email          `json:"email"`
	Casbin         Casbin         `json:"casbin"`
	System         System         `json:"system"`
	Captcha        Captcha        `json:"captcha"`
}
