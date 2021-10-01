package system

import "github.com/flipped-aurora/gf-vue-admin/library/global"

type Dictionary struct {
	global.Model
	Desc              string             `json:"desc" gorm:"column:desc;comment:描述"`
	Name              string             `json:"name" gorm:"column:name;comment:字典名(中)"`
	Type              string             `json:"type" gorm:"column:type;comment:字典名(英)"`
	Status            *bool              `json:"status" gorm:"column:status;comment:状态"`
	DictionaryDetails []DictionaryDetail `json:"sysDictionaryDetails"`
}

func (s *Dictionary) TableName() string {
	return "system_dictionaries"
}
