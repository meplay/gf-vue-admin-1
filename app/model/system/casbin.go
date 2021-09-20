package system

type Casbin struct {
	PType       string `gorm:"column:p_type"`
	AuthorityId string `gorm:"column:v0"`
	Path        string `gorm:"column:v1"`
	Method      string `gorm:"column:v2"`
}

func (c *Casbin) TableName() string {
	return "casbin_rule"
}
