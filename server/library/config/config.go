package config

type Config struct {
	Jwt            Jwt            `json:"jwt"`
	Oss            Oss            `json:"oss"`
	Redis          Redis          `json:"redis"`
	Email          Email          `json:"email"`
	Casbin         Casbin         `json:"casbin"`
	Logger         Logger         `json:"logger"`
	Server         Server         `json:"server"`
	System         System         `json:"system"`
	Captcha        Captcha        `json:"captcha"`
	Database       Database       `json:"database"`
	DatabaseLogger DatabaseLogger `json:"database_logger"`
}
