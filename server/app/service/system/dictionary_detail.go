package service

import (
	model "flipped-aurora/gf-vue-admin/server/app/model/system"
	"flipped-aurora/gf-vue-admin/server/app/model/system/request"
	"flipped-aurora/gf-vue-admin/server/app/service/system/internal"
	"flipped-aurora/gf-vue-admin/server/library/global"
)

var DictionaryDetail = new(detail)

type detail struct{
	_detail model.DictionaryDetail
}

// Create 创建字典详情数据
// Author [Aizen1172](https://github.com/Aizen1172)
func (d *detail) Create(info *request.CreateDictionaryDetail) error {
	return global.Db.Create(info.Create()).Error
}

// First 根据id获取字典详情单条数据
// Author [Aizen1172](https://github.com/Aizen1172)
func (d *detail) First(info *request.GetById) (result *model.DictionaryDetail, err error) {
	var entity model.DictionaryDetail
	err = global.Db.Where("id = ?",info.Id).First(&entity).Error
	return &entity, err
}

// Update 更新字典详情数据
// Author [Aizen1172](https://github.com/Aizen1172)
func (d *detail) Update(info *request.UpdateDictionaryDetail) error {
	return global.Db.Where("id = ?",info.Id).Updates(info.Update()).Error
}

// Delete 删除字典详情数据
// Author [Aizen1172](https://github.com/Aizen1172)
func (d *detail) Delete(info *request.GetById) error {
	return global.Db.Delete(&model.DictionaryDetail{}, info.Id).Error
}

// GetList 分页获取字典详情列表
// Author [Aizen1172](https://github.com/Aizen1172)
func (d *detail) GetList(info *request.SearchDictionaryDetail) (list interface{}, total int64, err error) {
	var details []model.DictionaryDetail
	db := global.Db.Model(&model.DictionaryDetail{})
	db.Scopes(info.Search())
	err = db.Count(&total).Error
	err = db.Scopes(internal.Gorm.Paginate(&info.PageInfo)).Find(&details).Error
	return details, total, err
}