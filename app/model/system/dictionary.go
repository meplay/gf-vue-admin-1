package system

import "github.com/flipped-aurora/gf-vue-admin/library/global"

type Dictionary struct {
	global.Model
	Name                 string             `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`   // 字典名（中）
	Type                 string             `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`   // 字典名（英）
	Status               *bool              `json:"status" form:"status" gorm:"column:status;comment:状态"` // 状态
	SysDictionaryDetails []DictionaryDetail `json:"sysDictionaryDetails" form:"sysDictionaryDetails"`
}

func (s *Dictionary) TableName() string {
	return "dictionary"
}
