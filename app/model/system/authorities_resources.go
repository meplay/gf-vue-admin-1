package system

// AuthoritiesResources 角色资源表
type AuthoritiesResources struct {
	AuthorityId string `gorm:"column:authority_id"`
	ResourcesId string `gorm:"column:resources_id"`
}

func (a *AuthoritiesResources) TableName() string {
	return "system_authorities_resources"
}
