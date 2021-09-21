package system

import (
	"gorm.io/gorm"
	"time"
)

type Authority struct {
	CreatedAt          time.Time      `json:"CreatedAt" gorm:"column:created_at;comment:创建时间" example:"创建时间"`
	UpdatedAt          time.Time      `json:"UpdatedAt" gorm:"column:updated_at;comment:更新时间" example:"更新时间"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index;column:deleted_at;comment:删除时间" example:"删除时间"`
	ParentId           string         `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	AuthorityId        string         `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName      string         `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	DefaultRouter      string         `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`                 // 默认菜单(默认dashboard)
	Menus              []Menu         `json:"menus" gorm:"many2many:authorities_menus;foreignKey:AuthorityId;joinForeignKey:AuthorityId;References:ID;JoinReferences:MenuID"`
	Children           []Authority    `json:"children" gorm:"-"`
	AuthorityResources []Authority    `json:"dataAuthorityId" gorm:"many2many:authorities_resources;foreignKey:AuthorityId;joinForeignKey:AuthorityId;References:AuthorityId;JoinReferences:ResourcesId"`
}

func (a *Authority) TableName() string {
	return "authorities"
}
