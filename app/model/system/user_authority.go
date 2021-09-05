package system

type UseAuthority struct {
	UserId      uint   `gorm:"column:user_id"`
	AuthorityId string `gorm:"column:authority_id"`
}

func (s *UseAuthority) TableName() string {
	return "users_authorities"
}
