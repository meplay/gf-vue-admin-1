package boot

import (
	"flag"
	"fmt"
	"gf-vue-admin/library/constant"
	"gf-vue-admin/library/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var Viper = new(_viper)

type _viper struct {
	err  error
	path string
}

func (v *_viper) Initialize(path ...string) {
	if len(path) == 0 {
		flag.StringVar(&v.path, "c", "", "choose config file.")
		flag.Parse()
		if v.path == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(constant.ConfigEnv); configEnv == "" {
				v.path = constant.ConfigFile
				fmt.Println(`您正在使用 config 的默认值,config的路径为: `, v.path)
			} else {
				v.path = constant.ConfigEnv
				fmt.Println(`您正在使用 GVA_CONFIG 环境变量,config的路径为: `, v.path)
			}
		} else {
			fmt.Println(`您正在使用命令行的 -c 参数传递的值,config的路径为: `, v.path)
		}
	} else {
		v.path = path[0]
		fmt.Println(`您正在使用func (v *_viper) Initialize()传递的值,config的路径为: `, v.path)
	}

	var _v = viper.New()
	_v.SetConfigFile(v.path)
	if v.err = _v.ReadInConfig(); v.err != nil {
		panic(fmt.Sprintf(`读取config.yaml文件失败, err: %v`, v.err))
	}
	_v.WatchConfig()

	_v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(`配置文件已修改并更新,文件为: `, e.Name)
		if v.err = _v.Unmarshal(&global.Config); v.err != nil {
			fmt.Println(v.err)
		}
	})
	if v.err = _v.Unmarshal(&global.Config); v.err != nil {
		fmt.Println(`Json 序列化数据失败, err :`, v.err)
	}
	global.Viper = _v
}
