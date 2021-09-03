package system

import "github.com/flipped-aurora/gf-vue-admin/library/global"

type Api struct {
	global.Model
	Path        string `json:"path" gorm:"comment:api路径"`
	Method      string `json:"method" gorm:"default:POST;comment:方法"`
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`
	Description string `json:"description" gorm:"comment:api中文描述"`
}

func (a *Api) TableName() string {
	return "apis"
}
