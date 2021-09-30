package config

import "time"

type Gorm struct {
	Dsn             Dsn           `mapstructure:"dsn" json:"dns" yaml:"dsn"`
	Config          string        `mapstructure:"config" json:"config" yaml:"config"`
	LogZap          bool          `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
	LogMode         string        `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	AutoMigrate     bool          `mapstructure:"auto-migrate" json:"autoMigrate" yaml:"auto-migrate"`
	MaxIdleConnes   int           `mapstructure:"max-idle-connes" json:"maxIdleConnes" yaml:"max-idle-connes"`
	MaxOpenConnes   int           `mapstructure:"max-open-connes" json:"maxOpenConnes" yaml:"max-open-connes"`
	ConnMaxLifetime time.Duration `mapstructure:"conn-max-lifetime" json:"connMaxLifetime" yaml:"conn-max-lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn-max-idle-time" json:"connMaxIdleTime" yaml:"conn-max-idle-time"`
}

// GetMaxIdleConnes 获取 MaxIdleConnes 值, 局部不为0则取局部, 为0则取全局
// Author: SliverHorn
func (g *Gorm) GetMaxIdleConnes() int {
	if g.Dsn.MaxIdleConnes == 0 {
		return g.MaxIdleConnes
	}
	return g.Dsn.MaxIdleConnes
}

// GetMaxOpenConnes 获取 MaxOpenConnes 值, 局部不为0则取局部, 为0则取全局
// Author: SliverHorn
func (g *Gorm) GetMaxOpenConnes() int {
	if g.Dsn.MaxOpenConnes == 0 {
		return g.MaxOpenConnes
	}
	return g.Dsn.MaxOpenConnes
}

// GetConnMaxLifetime 获取 ConnMaxLifetime 值, 局部不为0则取局部, 为0则取全局
// Author: SliverHorn
func (g *Gorm) GetConnMaxLifetime() time.Duration {
	if int64(g.Dsn.ConnMaxLifetime) == 0 {
		return g.ConnMaxLifetime
	}
	return g.Dsn.ConnMaxLifetime
}

// GetConnMaxIdleTime 获取 ConnMaxIdleTime 值, 局部不为0则取局部, 为0则取全局
// Author: SliverHorn
func (g *Gorm) GetConnMaxIdleTime() time.Duration {
	if int64(g.Dsn.ConnMaxIdleTime) == 0 {
		return g.ConnMaxIdleTime
	}
	return g.Dsn.ConnMaxIdleTime
}
