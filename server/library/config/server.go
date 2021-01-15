package config

type Logger struct {
	Path   string `json:"path"`
	Level  string `json:"level"`
	Stdout bool   `json:"stdout"`
}

type Server struct {
	LogPath          string `json:"log_path"`
	Address          string `json:"address"`
	DumpRouterMap    bool   `json:"dump_router_map"`
	ErrorLogEnabled  bool   `json:"error_log_enabled"`
	AccessLogEnabled bool   `json:"access_log_enabled"`
}