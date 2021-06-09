package internal

import (
	"gf-vue-admin/library/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Gorm = new(_gorm)

type _gorm struct{}

// GenerateConfig 根据配置文件生成对应 gorm.Config
// Author: SliverHorn
func (g *_gorm) GenerateConfig() *gorm.Config {
	_config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.Config.Gorm.LogZap {
	case "silent", "Silent":
		_config.Logger = Default.LogMode(logger.Silent)
	case "error", "Error":
		_config.Logger = Default.LogMode(logger.Error)
	case "warn", "Warn":
		_config.Logger = Default.LogMode(logger.Warn)
	case "info", "Info":
		_config.Logger = Default.LogMode(logger.Info)
	default:
		if global.Config.Gorm.LogMode {
			_config.Logger = Default.LogMode(logger.Info)
			break
		}
		_config.Logger = Default.LogMode(logger.Silent)
	}
	return _config
}
