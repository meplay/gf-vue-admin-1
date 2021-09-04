package request

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"gorm.io/gorm"
)

type DictionaryDetailCreate struct {
	Sort         int    `json:"sort" example:"排序标记"`
	Value        int    `json:"value" example:"字典值"`
	Label        string `json:"label" example:"展示值"`
	Status       *bool  `json:"status" example:"启用状态"`
	DictionaryID int    `json:"dictionaryId" example:"关联标记"`
}

func (r *DictionaryDetailCreate) Create() system.DictionaryDetail {
	return system.DictionaryDetail{Sort: r.Sort, Value: r.Value, Label: r.Label, Status: r.Status, DictionaryID: r.DictionaryID}
}

type DictionaryDetailUpdate struct {
	common.GetByID
	Sort         int    `json:"sort" example:"排序标记"`
	Value        int    `json:"value" example:"字典值"`
	Label        string `json:"label" example:"展示值"`
	Status       *bool  `json:"status" example:"启用状态"`
	DictionaryID int    `json:"dictionaryId" example:"关联标记"`
}

func (r *DictionaryDetailUpdate) Update() system.DictionaryDetail {
	return system.DictionaryDetail{Sort: r.Sort, Value: r.Value, Label: r.Label, Status: r.Status, DictionaryID: r.DictionaryID}
}

type DictionaryDetailSearch struct {
	common.PageInfo
	Sort         int    `json:"sort" example:"排序标记"`
	Value        int    `json:"value" example:"字典值"`
	Label        string `json:"label" example:"展示值"`
	Status       *bool  `json:"status" example:"启用状态"`
	DictionaryID int    `json:"dictionaryId" example:"关联标记"`
}

func (r *DictionaryDetailSearch) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Label != "" {
			db = db.Where("label LIKE ?", "%"+r.Label+"%")
		}
		if r.Value != 0 {
			db = db.Where("value = ?", r.Value)
		}
		if r.Status != nil {
			db = db.Where("status = ?", r.Status)
		}
		if r.DictionaryID != 0 {
			db = db.Where("dictionary_id = ?", r.DictionaryID)
		}
		return db
	}
}
