package service

import (
	"gf-vue-admin/app/api/request"
	"gf-vue-admin/app/api/response"
	model "gf-vue-admin/app/model/system"
	"github.com/gogf/gf/frame/g"
)

var Api = new(api)

type api struct {}

func (a *api) Create(info *request.CreateApi) error {
	if _, err := g.DB().Table("apis").FindOne(g.Map{"path": info.Path, "method": info.Method}); err != nil {
		return response.ErrorSameApi
	}
	entity := model.Api{
		Path:        info.Path,
		Method:      info.Method,
		ApiGroup:    info.ApiGroup,
		Description: info.Description,
	}
	if _, err := g.DB().Table("apis").Insert(&entity); err != nil {
		return err
	}
	return nil
}

func (a *api) Update()  {
	
}