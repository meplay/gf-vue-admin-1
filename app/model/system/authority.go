package system

import "time"

type Authority struct {
	CreatedAt          time.Time   // 创建时间
	UpdatedAt          time.Time   // 更新时间
	DeletedAt          *time.Time  `sql:"index"`
	AuthorityId        string      `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName      string      `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	ParentId           string      `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	DefaultRouter      string      `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`                 // 默认菜单(默认dashboard)
	Menus              []Menu      `json:"menus" gorm:"many2many:authorities_menus;foreignKey:AuthorityId;joinForeignKey:AuthorityId;References:ID;JoinReferences:MenuID"`
	Children           []Authority `json:"children" gorm:"-"`
	AuthorityResources []Authority `json:"dataAuthorityId" gorm:"many2many:authorities_resources;foreignKey:AuthorityId;joinForeignKey:AuthorityId;References:AuthorityId;JoinReferences:ResourcesId"`
}

func (a *Authority) TableName() string {
	return "authorities"
}
