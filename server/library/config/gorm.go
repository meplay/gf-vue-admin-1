package config

type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConnes int    `mapstructure:"max-idle-connes" json:"maxIdleConnes" yaml:"max-idle-connes"`
	MaxOpenConnes int    `mapstructure:"max-open-connes" json:"maxOpenConnes" yaml:"max-open-connes"`
	LogMode      bool `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap       string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
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

