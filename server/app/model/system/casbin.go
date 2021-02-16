package model

type Casbin struct {
	Path        string `json:"path" gorm:"column:v1"`
	PType       string `json:"ptype" gorm:"column:p_type"`
	Method      string `json:"method" gorm:"column:v2"`
	AuthorityId string `json:"rolename" gorm:"column:v0"`
}

func (c *Casbin) TableName() string {
	return "casbin_rule"
}

