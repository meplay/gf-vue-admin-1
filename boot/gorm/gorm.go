package boot

import (
	"log"
	"os"
	"time"

	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Gorm = new(_gorm)

type _gorm struct{}

// GenerateConfig 根据配置文件生成对应 *gorm.Config
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) GenerateConfig() *gorm.Config {
	_config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})
	switch global.Config.Gorm.LogMode {
	case "silent", "Silent":
		_config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		_config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		_config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		_config.Logger = _default.LogMode(logger.Info)
	default:
		_config.Logger = _default.LogMode(logger.Info)
	}
	_config.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   "",    // 表名前缀
		SingularTable: false, // 使用单数表名，启用该选项后，`User` 表将是`user`
		NameReplacer:  nil,   // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称
		NoLowerCase:   false, // 无小写
	}
	return _config
}
