package config

type System struct {
	Db            string `json:"db"`
	Env           string `json:"env"`
	OssType       string `json:"oss_type"`
	ErrorToEmail  bool   `json:"error_to_email"`
	UseMultipoint bool   `json:"use_multipoint"`
}