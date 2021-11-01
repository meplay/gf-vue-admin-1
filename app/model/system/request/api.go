package request

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"gorm.io/gorm"
)

type ApiCreate struct {
	Path        string `json:"path" example:"/api/create"`
	Method      string `json:"method" example:"请求方法:创建POST(默认)|查看GET|更新PUT|删除DELETE"`
	ApiGroup    string `json:"apiGroup" example:"api组"`
	Description string `json:"description" example:"api中文描述"`
}

func (r *ApiCreate) Create() system.Api {
	return system.Api{
		Path:        r.Path,
		Method:      r.Method,
		ApiGroup:    r.ApiGroup,
		Description: r.Description,
	}
}

type ApiUpdate struct {
	common.GetByID
	Path        string `json:"path" example:"api路径"`
	Method      string `json:"method" example:"api组"`
	ApiGroup    string `json:"apiGroup" example:"api请求方法"`
	Description string `json:"description"  example:"中文描述"`
}

func (r *ApiUpdate) Update() system.Api {
	return system.Api{Path: r.Path, Method: r.Method, ApiGroup: r.ApiGroup, Description: r.Description}
}

type ApiDelete struct {
	common.GetByID
	Path   string `json:"path" example:"/api/create"`
	Method string `json:"method" example:"请求方法:创建POST(默认)|查看GET|更新PUT|删除DELETE"`
}

type ApiSearch struct {
	common.PageInfo
	Desc        bool   `json:"desc" example:"false"`
	Path        string `json:"path" example:"/api/create"`
	Order       string `json:"orderKey" example:"排序"`
	Method      string `json:"method" example:"请求方法:创建POST(默认)|查看GET|更新PUT|删除DELETE"`
	ApiGroup    string `json:"apiGroup" example:"api组"`
	Description string `json:"description" example:"api中文描述"`
}

func (r *ApiSearch) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Path != "" {
			db = db.Where("path LIKE ?", "%"+r.Path+"%")
		}
		if r.Description != "" {
			db = db.Where("description LIKE ?", "%"+r.Description+"%")
		}
		if r.Method != "" {
			db = db.Where("method = ?", r.Method)
		}
		if r.ApiGroup != "" {
			db = db.Where("api_group = ?", r.ApiGroup)
		}
		if r.Order != "" {
			if r.Desc {
				db = db.Order(r.Order + " desc")
			} else {
				db = db.Order(r.Order)
			}
		} else {
			db = db.Order("api_group")
		}
		return db
	}
}
