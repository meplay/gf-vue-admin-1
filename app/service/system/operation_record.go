package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
)

var OperationRecord = new(operationRecord)

type operationRecord struct{}

// Create 创建记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *operationRecord) Create(info *request.OperationRecordCreate) error {
	return global.Db.Create(&info.OperationRecord).Error
}

// First 根据id获取单条操作记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *operationRecord) First(info *common.GetByID) (data *system.OperationRecord, err error) {
	var entity system.OperationRecord
	if err = global.Db.Where("id = ?", info.ID).First(&entity).Error; err != nil {
		return nil, errors.Wrap(err, "查找记录失败")
	}
	return &entity, nil
}

// Delete 删除操作记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *operationRecord) Delete(info *common.GetByID) error {
	return global.Db.Delete(&system.OperationRecord{}, info.ID).Error
}

// Deletes 批量删除记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *operationRecord) Deletes(ids *common.GetByIDs) error {
	return global.Db.Delete(&[]system.OperationRecord{}, "id in (?)", ids.Ids).Error
}

// GetList 分页获取操作记录列表
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *operationRecord) GetList(info *request.OperationRecordSearch) (list []system.OperationRecord, total int64, err error) {
	db := global.Db.Model(&system.OperationRecord{})
	var entities []system.OperationRecord
	err = db.Scopes(info.Search()).Count(&total).Preload("User").Find(&entities).Error
	return entities, total, err
}
