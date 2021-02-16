package model

type Casbin struct {
	Path        string `orm:"v1" json:"path" gorm:"column:v1"`
	PType       string `orm:"p_type" json:"ptype" gorm:"column:p_type"`
	Method      string `orm:"v2" json:"method" gorm:"column:v2"`
	AuthorityId string `orm:"v0" json:"rolename" gorm:"column:v0"`
}

func (c *Casbin) TableName() string {
	return "casbin_rule"
}

