package config

import "time"

type Database struct {
	Host             string        `json:"host"`
	Port             string        `json:"port"`
	User             string        `json:"user"`
	Pass             string        `json:"pass"`
	Name             string        `json:"name"`
	Type             string        `json:"type"`
	Role             string        `json:"role"`
	Debug            bool          `json:"debug"`
	Prefix           string        `json:"prefix"`
	DryRun           bool          `json:"dry_run"`
	Weight           int           `json:"weight"`
	Charset          string        `json:"charset"`
	LinkInfo         string        `json:"link_info"`
	MaxIdleConnCount int           `json:"max_idle_conn_count"`
	MaxOpenConnCount int           `json:"max_open_conn_count"`
	MaxConnLifetime  time.Duration `json:"max_conn_lifetime"`
}

type DatabaseLogger struct {
	Path   string `json:"path"`
	Level  string `json:"level"`
	Stdout bool   `json:"stdout"`
}
