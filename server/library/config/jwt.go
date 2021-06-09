package config

import "time"

type Jwt struct {
	ExpiresAt  time.Duration `mapstructure:"expires-at" json:"expiresAt" yaml:"expires-at"`
	RefreshAt  time.Duration `mapstructure:"refresh-at" json:"refreshAt" yaml:"refresh-at"`
	SigningKey string        `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
}
