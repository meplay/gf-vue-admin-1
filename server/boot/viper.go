package boot

import (
	"flag"
	"flipped-aurora/gf-vue-admin/server/library/constant"
	"flipped-aurora/gf-vue-admin/server/library/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gogf/gf/frame/g"
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
			if env := os.Getenv(constant.ConfigEnv); env == "" {
				v.path = constant.ConfigPostgresFile
				g.Log().Info(`您正在使用 config 的默认值!`, g.Map{"path": v.path})
			} else {
				switch env {
				case "mysql", "Mysql":
					v.path = constant.ConfigMysqlFile
					g.Log().Info(`您正在使用 GFVA_CONFIG环境变量 的 Mysql 环境 配置文件!`, g.Map{"path": v.path})
				case "postgres", "Postgres":
					v.path = constant.ConfigPostgresFile
					g.Log().Info(`您正在使用 GFVA_CONFIG环境变量 的 Postgres 环境 配置文件!`, g.Map{"path": v.path})
				case "develop", "Develop":
					v.path = constant.ConfigDevelopFile
					g.Log().Info(`您正在使用 GFVA_CONFIG环境变量 的 Develop 环境 配置文件!`, g.Map{"path": v.path})
				case "production", "Production":
					v.path = constant.ConfigProductionFile
					g.Log().Info(`您正在使用 GFVA_CONFIG环境变量 的 Production 环境 配置文件!`, g.Map{"path": v.path})
				case "docker-compose", "DockerCompose", "Docker-Compose":
					v.path = constant.ConfigDockerComposeFile
					g.Log().Info(`您正在使用 GFVA_CONFIG环境变量 的 Docker Compose 环境 配置文件!`, g.Map{"path": v.path})
				}
				v.path = constant.ConfigEnv
				g.Log().Info(`您正在使用 GFVA_CONFIG 环境变量!`, g.Map{"path": v.path})
			}
		} else {
			g.Log().Info(`您正在使用命令行的 -c 参数传递的值!`, g.Map{"path": v.path})
		}
	} else {
		v.path = path[0]
		g.Log().Info(`您正在使用func (v *_viper) Initialize()传递的值!`, g.Map{"path": v.path})
	}

	_v := viper.New()
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
