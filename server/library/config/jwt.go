package config

type Jwt struct {
	ExpiresAt  int    `json:"expires_at"`
	RefreshAt  int    `json:"refresh_at"`
	SigningKey string `json:"signing_key"`
}
