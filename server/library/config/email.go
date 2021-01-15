package config

type Email struct {
	To       string `json:"to"`
	Port     int    `json:"port"`
	From     string `json:"from"`
	Host     string `json:"host"`
	IsSsl    bool   `json:"is_ssl"`
	Secret   string `json:"secret"`
	Nickname string `json:"nickname"`
}
