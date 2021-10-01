package system

// AuthoritiesMenus 角色按钮关联表
type AuthoritiesMenus struct {
	MenuId      uint   `json:"menuId" gorm:"column:menu_id" swaggertype:"string" example:"uint 菜单ID"`
	AuthorityId string `json:"authorityId" gorm:"column:authority_id" example:"角色ID"`
}

func (a *AuthoritiesMenus) TableName() string {
	return "system_authorities_menus"
}
