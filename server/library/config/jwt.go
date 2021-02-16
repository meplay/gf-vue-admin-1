package config

type Jwt struct {
	ExpiresAt  int    `mapstructure:"expires-at" json:"expiresAt" yaml:"expires-at"`
	RefreshAt  int    `mapstructure:"refresh-at" json:"refreshAt" yaml:"refresh-at"`
	SigningKey string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
}
