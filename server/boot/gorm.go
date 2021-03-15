package boot

import (
	"gf-vue-admin/library/global"
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
	if global.GormConfig.Path != "" {
		global.Config.Mysql = global.GormConfig
	} else {
		global.Config.Mysql = global.Config.Mysql.GetByLink()
	}

}
