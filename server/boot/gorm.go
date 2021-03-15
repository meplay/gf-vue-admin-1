package boot

import (
	"gf-vue-admin/library/config"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
	"strings"
)

var Gorm = new(_gorm)

type _gorm struct{}

func (g *_gorm) Initialize() {
	switch global.Config.System.DbType {
	case "mysql":
		Mysql.Initialize()
	}
}

func init() {
	link := g.Cfg().GetString("database.default.link")
	a := strings.Split(link, ":")
	if len(a) < 4 {
		g.Log().Error("获取失败!")
	}
	user := a[1]
	passAndHost := strings.Split(a[2], "@")
	portAndName := strings.Split(a[3], "/")
	global.Config.Mysql = config.Mysql{
		Path:         strings.Split(passAndHost[1], "(")[1] + ":" + strings.Split(strings.Split(portAndName[0], "/")[0], ")")[0],
		Config:        "charset=utf8mb4&parseTime=True&loc=Local",
		Dbname:        portAndName[1],
		Username:      user,
		Password:      passAndHost[0],
		MaxIdleConnes: 10,
		MaxOpenConnes: 10,
		LogMode:       false,
		LogZap:        "",
	}
}