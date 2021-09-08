package config

type Jwt struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // 签名
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // 缓冲时间
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
}
