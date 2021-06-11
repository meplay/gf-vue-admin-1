package service

import (
	model "flipped-aurora/gf-vue-admin/server/app/model/system"
	"flipped-aurora/gf-vue-admin/server/app/model/system/request"
	"flipped-aurora/gf-vue-admin/server/app/service/system/internal"
	"flipped-aurora/gf-vue-admin/server/library/global"
)

var OperationRecord = new(record)

type record struct {
	_admin  model.Admin
	_record model.OperationRecord
}

// Create 创建记录
// Author [Aizen1172](https://github.com/Aizen1172)
func (r *record) Create(info *request.CreateOperationRecord) error {
	return global.Db.Create(info.Create()).Error
}

// First 根据id获取单条操作记录
// Author [Aizen1172](https://github.com/Aizen1172)
func (r *record) First(info *request.GetById) (result *model.OperationRecord, err error) {
	var entity model.OperationRecord
	err = global.Db.Where("id = ?", info.Id).First(&entity).Error
	return &entity, err
}

// Delete 删除操作记录
// Author [Aizen1172](https://github.com/Aizen1172)
func (r *record) Delete(info *request.GetById) error {
	return global.Db.Delete(&model.OperationRecord{}, info.Id).Error
}

// Deletes 批量删除记录
// Author [Aizen1172](https://github.com/Aizen1172)
func (r *record) Deletes(info *request.GetByIds) error {
	return global.Db.Delete(&model.OperationRecord{}, info.Ids).Error
}

// GetList 分页获取操作记录列表
// Author [Aizen1172](https://github.com/Aizen1172)
func (r *record) GetList(info *request.SearchOperationRecord) (list interface{}, total int64, err error) {
	var records []model.OperationRecord
	db := global.Db.Model(&model.OperationRecord{})
	db.Scopes(info.Search())
	err = db.Count(&total).Error
	err = db.Scopes(internal.Gorm.Paginate(&info.PageInfo)).Find(&records).Error
	for i, operation := range records {
		err = global.Db.Where("id = ?", operation.UserID).First(&records[i].Admin).Error
	}
	return records, total, err
}
