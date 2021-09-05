package system

import "github.com/flipped-aurora/gf-vue-admin/library/global"

type Api struct {
	global.Model
	Path        string `json:"path" gorm:"column:path;comment:api路径"`
	Method      string `json:"method" gorm:"column:method;default:POST;comment:方法"`
	ApiGroup    string `json:"apiGroup" gorm:"column:api_group;comment:api组"`
	Description string `json:"description" gorm:"column:description;comment:api中文描述"`
}

func (a *Api) TableName() string {
	return "apis"
}
