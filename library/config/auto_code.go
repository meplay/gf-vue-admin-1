package config

type AutoCode struct {
	Web Web `mapstructure:"web" json:"web" yaml:"web"`

	Root string `mapstructure:"root" json:"root" yaml:"root"`

	Server Server `mapstructure:"server" json:"server" yaml:"server"`

	Restart bool `mapstructure:"restart" json:"restart" yaml:"restart"`

	RubbishPath string `mapstructure:"rubbish-path" json:"rubbishPath" yaml:"rubbish-path"`
}

type Server struct {
	Api     string `mapstructure:"api" json:"api" yaml:"api"`
	Boot    string `mapstructure:"boot" json:"boot" yaml:"boot"`
	Root    string `mapstructure:"root" json:"root" yaml:"root"`
	Model   string `mapstructure:"model" json:"model" yaml:"model"`
	Router  string `mapstructure:"router" json:"router" yaml:"router"`
	Request string `mapstructure:"request" json:"request" yaml:"request"`
	Service string `mapstructure:"service" json:"service" yaml:"service"`
}

type Web struct {
	Api   string `mapstructure:"api" json:"api" yaml:"api"`
	Root  string `mapstructure:"root" json:"root" yaml:"root"`
	Form  string `mapstructure:"form" json:"form" yaml:"form"`
	Flow  string `mapstructure:"flow" json:"flow" yaml:"flow"`
	Table string `mapstructure:"table" json:"table" yaml:"table"`
}
