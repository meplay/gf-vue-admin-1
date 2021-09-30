package system

import "github.com/flipped-aurora/gf-vue-admin/library/global"

type MenuParameter struct {
	global.Model
	MenuID uint   `json:"menu_id,omitempty" gorm:"column:menu_id;comment:menu ID"` // menu ID
	Key    string `json:"key" gorm:"comment:地址栏携带参数的key"`                          // 地址栏携带参数的key
	Type   string `json:"type" gorm:"comment:地址栏携带参数为params还是query"`               // 地址栏携带参数为params还是query
	Value  string `json:"value" gorm:"comment:地址栏携带参数的值"`                          // 地址栏携带参数的值
}

func (m *MenuParameter) TableName() string {
	return "menu_parameters"
}
