package service

import (
	model "flipped-aurora/gf-vue-admin/server/app/model/system"
	"flipped-aurora/gf-vue-admin/server/app/model/system/request"
	"flipped-aurora/gf-vue-admin/server/app/service/system/internal"
	"flipped-aurora/gf-vue-admin/server/library/global"
)

var Dictionary = new(dictionary)

type dictionary struct {
	_detail     model.DictionaryDetail
	_dictionary model.Dictionary
}

// Create 创建字典数据
// Author [Aizen1172](https://github.com/Aizen1172)
func (d *dictionary) Create(info *request.CreateDictionary) error {
	return global.Db.Create(info.Create()).Error
}

// First 根据id或者type获取字典单条数据
// Author [Aizen1172](https://github.com/Aizen1172)
func (d *dictionary) First(info *request.FirstDictionary) (result *model.Dictionary, err error) {
	var entity model.Dictionary
	err = global.Db.Where("id = ? OR type = ?", info.Id, info.Type).First(&entity).Error
	err = global.Db.Where("dictionary_id = ?", entity.ID).Find(&entity.DictionaryDetails).Error
	return &entity, err
}

// Update 更新字典数据
// Author [Aizen1172](https://github.com/Aizen1172)
func (d *dictionary) Update(info *request.UpdateDictionary) error {
	return global.Db.Where("id = ?",info.Id).Updates(info.Update()).Error
}

// Delete 删除字典数据
// Author [Aizen1172](https://github.com/Aizen1172)
func (d *dictionary) Delete(info *request.GetById) error {
	return global.Db.Delete(&model.Dictionary{}, info.Id).Error
}

// GetList 分页获取字典列表
// Author [Aizen1172](https://github.com/Aizen1172)
func (d *dictionary) GetList(info *request.SearchDictionary) (list interface{}, total int64, err error) {
	var dictionaries []model.Dictionary
	db := global.Db.Model(&model.Dictionary{})
	db.Scopes(info.Search())
	err = db.Count(&total).Error
	err = db.Scopes(internal.Gorm.Paginate(&info.PageInfo)).Find(&dictionaries).Error
	return dictionaries, total, err
}
