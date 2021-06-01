package service

import (
	"errors"
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/app/model/system/request"
	"gf-vue-admin/app/service/system/internal"
	"gf-vue-admin/library/global"
	"gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"gorm.io/gorm"
)

var Api = new(api)

type api struct {}

// Create 新增基础Api
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *api) Create(info *request.CreateApi) error {
	err := global.Db.Where("path = ? AND method = ?", info.Path, info.Method).First(&model.Api{}).Error
	if !errors.Is(err,gorm.ErrRecordNotFound) {
		g.Log().Error(response.ErrorSameApi, g.Map{"path": info.Path, "method": info.Method})
		return err
	}
	entity := info.Create()
	return global.Db.Create(&entity).Error
}

// First 根据id获取api
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *api) First(info *request.GetById) (result *model.Api, err error) {
	var entity model.Api
	err = global.Db.First(&entity,info.Id).Error
	return &entity, err
}

// Update 根据id更新api
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *api) Update(info *request.UpdateApi) error {
	var old *model.Api
	if err := global.Db.Where("id = ?",info.Id).First(&old).Error;err != nil {
		return err
	}
	if old.Path == info.Path || old.Method == info.Method {
		return response.ErrorSameApi
	}
	err := global.Db.Where("id = ?",info.Id).Updates(info.Update()).Error
	err = Casbin.UpdateApi(old.Path, info.Path, old.Method, info.Method)
	return err
}

// Delete 删除基础api
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *api) Delete(info *request.DeleteApi) error {
	Casbin.Clear(info.Path, info.Method)
	return global.Db.Delete(&model.Api{}, info.Id).Error
}

// Deletes 删除选中API
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *api) Deletes(info *request.GetByIds) error {
	return global.Db.Delete(&model.Api{}, info.Ids).Error
}

// GetList 分页获取数据
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *api) GetList(info *request.SearchApi) (list interface{}, total int64, err error) {
	var apis []model.Api
	db := global.Db.Model(&model.Api{})
	db.Scopes(info.Search())
	err = db.Count(&total).Error
	err = db.Scopes(internal.Gorm.Paginate(&info.PageInfo)).Find(&apis).Error
	return apis, total, err
}

// GetAllApi 获取所有api
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *api) GetAllApi() (result *[]model.Api, err error) {
	var apis []model.Api
	err = global.Db.Find(&apis).Error
	return &apis, err
}
