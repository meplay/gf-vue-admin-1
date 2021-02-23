package config

import "fmt"

type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	From     string `mapstructure:"from" json:"from" yaml:"from"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	IsSsl    bool   `mapstructure:"is-ssl" json:"isSsl" yaml:"is-ssl"`
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"`
}

func (e *Email) GetFrom() string {
	if e.Nickname != "" {
		return fmt.Sprintf("%s <%s>", e.Nickname, e.From)
	} else {
		return e.From
	}
}
