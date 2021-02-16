package model

import (
	"gorm.io/gorm"
	"time"
)

type Authority struct {
	CreatedAt time.Time      `orm:"created_at" json:"CreatedAt"`
	UpdatedAt time.Time      `orm:"updated_at" json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `orm:"deleted_at" json:"-" gorm:"index"`

	ParentId      string `json:"parentId" gorm:"comment:父角色ID"`
	AuthorityId   string `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName string `json:"authorityName" gorm:"comment:角色名"`
	DefaultRouter string `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`

	Menus         []Menu      `orm:"-" json:"menus" gorm:"many2many:authorities_menus;foreignKey:AuthorityId;joinForeignKey:AuthorityId;References:ID;JoinReferences:MenuID"`
	Children      []Authority `orm:"-" json:"children" gorm:"-"`
	DataAuthority []Authority `orm:"-" json:"dataAuthorityId" gorm:"many2many:data_authorities;foreignKey:AuthorityId;joinForeignKey:AuthorityId;References:AuthorityId;JoinReferences:DataAuthority"`
}

func (a *Authority) TableName() string {
	return "authorities"
}

type DataAuthorities struct {
	AuthorityId   string `orm:"authority_id" gorm:"comment:角色id"`
	DataAuthority string `orm:"data_authority" gorm:"comment:资源id"`
}

func (d *DataAuthorities) TableName() string {
	return "data_authorities"
}

type AuthoritiesMenus struct {
	MenuId      uint
	AuthorityId string
}

func (a *AuthoritiesMenus) TableName() string {
	return "authorities_menus"
}

