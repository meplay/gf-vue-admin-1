package global

import (
	"github.com/flipped-aurora/gf-vue-admin/library/config"
	"gorm.io/gorm"
)

var (
	Db     *gorm.DB
	Config config.Config
)
