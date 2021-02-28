package constant

const (
	ConfigEnv      = "GVA_CONFIG"
	ConfigFile     = "config/viper.yaml"
	SerializerGob  = "gob"
	SerializerJson = "json"

	BasePath = "./template"
	AutoPath = "generate/"

	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB

	ExcelDir      = "./public/excel"
	FinishDir     = "./fileDir/"
	BreakpointDir = "./breakpointDir/"
)
