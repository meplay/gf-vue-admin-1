package system

import (
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Api struct {
	global.Model
	Path        string `json:"path" gorm:"column:path;comment:api路径"`
	Method      string `json:"method" gorm:"column:method;default:POST;comment:方法"`
	ApiGroup    string `json:"apiGroup" gorm:"column:api_group;comment:api组"`
	Description string `json:"description" gorm:"column:description;comment:api中文描述"`
}

func (a *Api) TableName() string {
	return "system_apis"
}

// BeforeCreate api创建前钩子函数
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *Api) BeforeCreate(tx *gorm.DB) error {
	if errors.Is(tx.Where("path = ? AND method = ?", a.Path, a.Method).First(&Api{}).Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.New("存在相同api")
}
