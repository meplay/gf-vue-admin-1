package system

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Casbin struct {
	PType       string `gorm:"column:p_type"`
	AuthorityId string `gorm:"column:v0"`
	Path        string `gorm:"column:v1"`
	Method      string `gorm:"column:v2"`
}

func (c *Casbin) BeforeCreate(tx *gorm.DB) error {
	entity := Casbin{}
	if errors.Is(tx.First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.Errorf(`角色id(%s:%s)存在相同api(%s:%s)!`, c.PType, c.AuthorityId, c.Path, c.Method)
}

func (c *Casbin) TableName() string {
	return "casbin_rule"
}
