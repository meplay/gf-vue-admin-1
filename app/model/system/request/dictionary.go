package request

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/gogf/gf/frame/g"
	"gorm.io/gorm"
)

type DictionaryCreate struct {
	Desc   string `json:"desc" example:"描述"`
	Name   string `json:"name" example:"字典名(中)"`
	Type   string `json:"type" example:"字典名(英)"`
	Status *bool  `json:"status" example:"状态"`
}

func (c *DictionaryCreate) Create() system.Dictionary {
	return system.Dictionary{Desc: c.Desc, Name: c.Name, Type: c.Type, Status: c.Status}
}

type DictionaryFirst struct {
	common.GetByID
	Type string `json:"type" example:"字典名(英)"`
}

type DictionaryUpdate struct {
	common.GetByID
	Desc   string `json:"desc" example:"描述"`
	Name   string `json:"name" example:"字典名(中)"`
	Type   string `json:"type" example:"字典名(英)"`
	Status *bool  `json:"status" example:"状态"`
}

func (u *DictionaryUpdate) Update() g.Map {
	return g.Map{"name": u.Name, "type": u.Type, "desc": u.Desc, "status": u.Status}
}

type DictionarySearch struct {
	common.PageInfo
	Desc   string `json:"desc" example:"描述"`
	Name   string `json:"name" example:"字典名(中)"`
	Type   string `json:"type" example:"字典名(英)"`
	Status *bool  `json:"status" example:"状态"`
}

func (d *DictionarySearch) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if d.Desc != "" {
			db = db.Where("`desc` LIKE ?", "%"+d.Desc+"%")
		}
		if d.Name != "" {
			db = db.Where("`name` LIKE ?", "%"+d.Name+"%")
		}
		if d.Type != "" {
			db = db.Where("`type` LIKE ?", "%"+d.Type+"%")
		}
		if d.Status != nil {
			db = db.Where("`status` = ?", d.Status)
		}
		return db
	}
}
