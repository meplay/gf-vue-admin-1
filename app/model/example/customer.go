package example

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
)

type Customer struct {
	global.Model
	UserID          uint        `json:"sysUserId" gorm:"comment:管理ID"`
	Name            string      `json:"customerName" gorm:"comment:客户名"`
	Phone           string      `json:"customerPhoneData" gorm:"comment:客户手机号码"`
	UserAuthorityID string      `json:"sysUserAuthorityID" gorm:"comment:管理角色ID"`
	User            system.User `json:"sysUser" gorm:"comment:管理详情"`
}

func (c *Customer) TableName() string {
	return "example_customer"
}
