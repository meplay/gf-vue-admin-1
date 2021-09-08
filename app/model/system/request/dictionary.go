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
	Status *bool  `json:"status"`
}

func (r *DictionaryCreate) Create() system.Dictionary {
	return system.Dictionary{Desc: r.Desc, Name: r.Name, Type: r.Type, Status: r.Status}
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
	Status *bool  `json:"status"`
}

func (r *DictionaryUpdate) Update() g.Map {
	return g.Map{"name": r.Name, "type": r.Type, "desc": r.Desc, "status": r.Status}
}

type DictionarySearch struct {
	common.PageInfo
	Desc   string `json:"desc" example:"描述"`
	Name   string `json:"name" example:"字典名(中)"`
	Type   string `json:"type" example:"字典名(英)"`
	Status *bool  `json:"status"`
}

func (r *DictionarySearch) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Desc != "" {
			db = db.Where("`desc` LIKE ?", "%"+r.Desc+"%")
		}
		if r.Name != "" {
			db = db.Where("`name` LIKE ?", "%"+r.Name+"%")
		}
		if r.Type != "" {
			db = db.Where("`type` LIKE ?", "%"+r.Type+"%")
		}
		if r.Status != nil {
			db = db.Where("`status` = ?", r.Status)
		}
		return db
	}
}
