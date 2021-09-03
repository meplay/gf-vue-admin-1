package request

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"gorm.io/gorm"
)

type CreateApi struct {
	Path        string `json:"path" example:"/api/create"`
	Method      string `json:"method" example:"请求方法:创建POST(默认)|查看GET|更新PUT|删除DELETE"`
	ApiGroup    string `json:"apiGroup" example:"api组"`
	Description string `json:"description" example:"api中文描述"`
}

type UpdateApi struct {
	common.GetByID
	Path        string `json:"path" example:"api路径"`
	Method      string `json:"method" example:"api组"`
	ApiGroup    string `json:"apiGroup" example:"api请求方法"`
	Description string `json:"description"  example:"中文描述"`
}

func (u *UpdateApi) Update() system.Api {
	return system.Api{Path: u.Path, Method: u.Method, ApiGroup: u.ApiGroup, Description: u.Description}

}

type DeleteApi struct {
	common.GetByID
	Path   string `json:"path" example:"/api/create"`
	Method string `json:"method" example:"请求方法:创建POST(默认)|查看GET|更新PUT|删除DELETE"`
}

type SearchApi struct {
	common.PageInfo
	Desc        bool   `json:"desc" example:"false"`
	Path        string `json:"path" example:"/api/create"`
	Order       string `json:"orderKey" example:"排序"`
	Method      string `json:"method" example:"请求方法:创建POST(默认)|查看GET|更新PUT|删除DELETE"`
	ApiGroup    string `json:"apiGroup" example:"api组"`
	Description string `json:"description" example:"api中文描述"`
}

func (s *SearchApi) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.Path != "" {
			db = db.Where("path LIKE ?", "%"+s.Path+"%")
		}
		if s.Description != "" {
			db = db.Where("description LIKE ?", "%"+s.Description+"%")
		}
		if s.Method != "" {
			db = db.Where("method = ?", s.Method)
		}
		if s.ApiGroup != "" {
			db = db.Where("api_group = ?", s.ApiGroup)
		}
		if s.Order != "" {
			if s.Desc {
				db = db.Order(s.Order + " desc")
			} else {
				db = db.Order(s.Order)
			}
		} else {
			db = db.Order("api_group")
		}
		return db
	}
}
