package config

import (
	"github.com/gogf/gf/frame/g"
	"strings"
)

type Mysql struct {
	Path          string `mapstructure:"path" json:"path" yaml:"path"`
	Config        string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname        string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username      string `mapstructure:"username" json:"username" yaml:"username"`
	Password      string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConnes int    `mapstructure:"max-idle-connes" json:"maxIdleConnes" yaml:"max-idle-connes"`
	MaxOpenConnes int    `mapstructure:"max-open-connes" json:"maxOpenConnes" yaml:"max-open-connes"`
	LogMode       bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap        string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetMaxIdleConnes() int {
	return m.MaxIdleConnes
}

func (m *Mysql) GetMaxOpenConnes() int {
	return m.MaxOpenConnes
}

func (m *Mysql) GetByLink() Mysql {
	var result Mysql
	// link = "mysql:root:gdkid,,..@tcp(127.0.0.1:13307)/gf_vue_admin
	link := g.Cfg().GetString("database.default.link")
	// a := []string{"mysql", "root", "gdkid,,..@tcp(127.0.0.1", "13307)/gf_vue_admin"}
	// a[0] = "mysql"
	// a[1] = "root
	// a[2] = "gdkid,,..@tcp(127.0.0.1"
	// a[3] = "13307)/gf_vue_admin"
	a := strings.Split(link, ":")
	if len(a) == 4 {
		result.Username = a[1] // root
		// b := []string{"gdkid,,..", "127.0.0.1"}
		b := strings.Split(a[2], "@tcp(")
		// c := []string{"13307", "gf_vue_admin"}
		c := strings.Split(a[3], ")/")
		if len(b) == 2 || len(c) == 2 {
			result.Password = b[0]          // gdkid,,..
			result.Path = b[1] + ":" + c[0] // 127.0.0.1:13307
			result.Dbname = c[1]
		}
		result.Config = "charset=utf8mb4&parseTime=True&loc=Local"
		result.LogZap = ""
		result.LogMode = false
		result.MaxIdleConnes = 10
		result.MaxOpenConnes = 100
		return result
	}
	return result

}
