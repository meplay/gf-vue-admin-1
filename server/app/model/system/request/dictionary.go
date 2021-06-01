package request

import (
	model "gf-vue-admin/app/model/system"
	"gorm.io/gorm"
)

type BaseDictionary struct {
	Name   string `p:"name" v:"required|length:1,1000#请输入字典中文名|字典中文名长度为:min到:max位"`
	Type   string `p:"type" v:"required|length:1,1000#请输入字典英文名|字典英文名长度为:min到:max位"`
	Desc   string `p:"desc" v:"required|length:1,1000#请输入描述|描述长度为:min到:max位"`
	Status *bool  `p:"status" v:"required|length:1,1000#请输入状态|状态长度为:min到:max位"`
}

type CreateDictionary struct {
	BaseDictionary
}

func (c *CreateDictionary) Create() *model.Dictionary {
	return &model.Dictionary{
		Name:              c.Name,
		Type:              c.Type,
		Status:            c.Status,
		Desc:              c.Desc,
		DictionaryDetails: nil,
	}
}

type UpdateDictionary struct {
	GetById
	BaseDictionary
}

func (u *UpdateDictionary) Update() *model.Dictionary {
	return &model.Dictionary{
		Name:              u.Name,
		Type:              u.Type,
		Status:            u.Status,
		Desc:              u.Desc,
	}
}

type GetDictionary struct {
	ID   float64 `p:"id" v:"required|length:1,1000#请输入id|id长度为:min到:max位"`
	Type string  `p:"type" v:"required|length:1,1000#请输入字典英文名|字典英文名长度为:min到:max位"`
}

// FindDictionary 用id查询Dictionary
type FirstDictionary struct {
	Id   int `p:"id" v:"required|length:1,1000#请输入id|id长度为:min到:max位"` // 自增ID
	Type string  `p:"type" `
}

type SearchDictionary struct {
	Status *bool  `p:"status"`
	Name   string `p:"name"`
	Type   string `p:"type"`
	Desc   string `p:"desc"`
	PageInfo
}

func (s *SearchDictionary) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.Name != "" {
			db.Where("name LIKE ?","%" + s.Name + "%")
		}
		if s.Type != "" {
			db.Where("type LIKE ?","%" + s.Type + "%")
		}
		if s.Status != nil {
			if *s.Status == true {
				db.Where("status = ?",1)
			} else {
				db.Where("status = ?",2)
			}
		}
		if s.Desc != "" {
			db.Where("desc LIKE ?","%" + s.Desc + "%")
		}
		return db
	}
}
