package config

import (
	"time"
)

type Gorm struct {
	Dsn             GormConfig    `mapstructure:"dsn" json:"dns" yaml:"dsn"`
	Config          string        `mapstructure:"config" json:"config" yaml:"config"`
	LogZap          string        `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
	LogMode         bool          `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	AutoMigrate     bool          `mapstructure:"auto-migrate" json:"autoMigrate" yaml:"auto-migrate"`
	MaxIdleConnes   int           `mapstructure:"max-idle-connes" json:"maxIdleConnes" yaml:"max-idle-connes"`
	MaxOpenConnes   int           `mapstructure:"max-open-connes" json:"maxOpenConnes" yaml:"max-open-connes"`
	ConnMaxLifetime time.Duration `mapstructure:"conn-max-lifetime" json:"connMaxLifetime" yaml:"conn-max-lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn-max-idle-time" json:"connMaxIdleTime" yaml:"conn-max-idle-time"`
}

type GormConfig struct {
	MaxIdleConnes   int           `mapstructure:"max-idle-connes" json:"maxIdleConnes" yaml:"max-idle-connes"`
	MaxOpenConnes   int           `mapstructure:"max-open-connes" json:"maxOpenConnes" yaml:"max-open-connes"`
	ConnMaxLifetime time.Duration `mapstructure:"conn-max-lifetime" json:"connMaxLifetime" yaml:"conn-max-lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn-max-idle-time" json:"connMaxIdleTime" yaml:"conn-max-idle-time"`
	Sources         []Source      `mapstructure:"sources" json:"sources" yaml:"sources"`
	Replicas        []Replica     `mapstructure:"replicas" json:"replicas" yaml:"replicas"`
}

type Source struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	Port        string `mapstructure:"port" json:"port" yaml:"port"`
	DbName      string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username    string `mapstructure:"username" json:"username" yaml:"username"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	OtherConfig string `mapstructure:"other-config" json:"otherConfig" yaml:"other-config"`
}

func (s *Source) GetMysqlDsn(config string) string {
	if s.OtherConfig != "" {
		return s.Username + ":" + s.Password + "@tcp(" + s.Host + ":" + s.Port + ")/" + s.DbName + "?" + s.OtherConfig
	}
	return s.Username + ":" + s.Password + "@tcp(" + s.Host + ":" + s.Port + ")/" + s.DbName + "?" + config
}

func (s *Source) GetPostgresDsn(config string) string {
	if s.OtherConfig != "" {
		return "host=" + s.Host + " user=" + s.Username + " password=" + s.Password + " dbname=" + s.DbName + " port=" + s.Port + " " + s.OtherConfig
	}
	return "host=" + s.Host + " user=" + s.Username + " password=" + s.Password + " dbname=" + s.DbName + " port=" + s.Port + " " + config
}

type Replica struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	Port        string `mapstructure:"port" json:"port" yaml:"port"`
	DbName      string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username    string `mapstructure:"username" json:"username" yaml:"username"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	OtherConfig string `mapstructure:"other-config" json:"otherConfig" yaml:"other-config"`
}

func (r *Replica) GetMysqlDsn(config string) string {
	if r.OtherConfig != "" {
		return r.Username + ":" + r.Password + "@tcp(" + r.Host + ":" + r.Port + ")/" + r.DbName + "?" + r.OtherConfig
	}
	return r.Username + ":" + r.Password + "@tcp(" + r.Host + ":" + r.Port + ")/" + r.DbName + "?" + config
}

func (r *Replica) GetPostgresDsn(config string) string {
	if r.OtherConfig != "" {
		return "host=" + r.Host + " user=" + r.Username + " password=" + r.Password + " dbname=" + r.DbName + " port=" + r.Port + " " + r.OtherConfig
	}
	return "host=" + r.Host + " user=" + r.Username + " password=" + r.Password + " dbname=" + r.DbName + " port=" + r.Port + " " + config
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
