package service

import (
	"gf-vue-admin/app/api/request"
	"gf-vue-admin/app/api/response"
	model "gf-vue-admin/app/model/system"
	"github.com/gogf/gf/frame/g"
)

var Api = new(api)

type api struct{}

func (a *api) Create(info *request.CreateApi) error {
	if _, err := g.DB().Table("apis").FindOne(g.Map{"path": info.Path, "method": info.Method}); err != nil {
		return response.ErrorSameApi
	}
	entity := info.Create()
	if _, err := g.DB().Table("apis").Insert(entity); err != nil {
		return err
	}
	return nil
}

func (a *api) Update(info *request.UpdateApi) error {
	var old model.Api
	if err := g.DB().Table("apis").Where(g.Map{"id": info.Id}).Struct(&old); err != nil {
		return err
	}
	if old.Path == info.Path || old.Method == info.Method {
		return response.ErrorSameApi
	}
	_, err := g.DB().Table("apis").Data(info.Update()).Update(g.Map{"id": info.Id})
	return err
}