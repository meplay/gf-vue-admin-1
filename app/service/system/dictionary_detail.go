package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	_errors "github.com/pkg/errors"
)

var DictionaryDetail = new(dictionaryDetail)

type dictionaryDetail struct{}

// Create 创建字典详情数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionaryDetail) Create(info *request.DictionaryDetailCreate) error {
	entity := info.Create()
	return global.Db.Create(&entity).Error
}

// First 根据id获取字典详情单条数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionaryDetail) First(info *common.GetByID) (data *system.DictionaryDetail, err error) {
	var entity system.DictionaryDetail
	if err = global.Db.Where("id = ?", info.ID).First(&entity).Error; err != nil {
		return nil, _errors.Wrap(err, "更新字典详情失败!")
	}
	return nil, err
}

// Update 更新字典详情数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionaryDetail) Update(info *request.DictionaryDetailUpdate) error {
	entity := info.Update()
	if err := global.Db.Model(&system.DictionaryDetail{}).Where("id = ?", info.ID).Updates(&entity).Error; err != nil {
		return _errors.Wrap(err, "更新字典详情失败!")
	}
	return nil
}

// Delete 删除字典详情数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionaryDetail) Delete(info *common.GetByID) error {
	return global.Db.First(&system.DictionaryDetail{}, info.ID).Error
}

// GetList 分页获取字典详情列表
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionaryDetail) GetList(info *request.DictionaryDetailSearch) (list []system.DictionaryDetail, total int64, err error) {
	entities := make([]system.DictionaryDetail, 0, info.PageSize)
	db := global.Db.Model(&system.DictionaryDetail{})
	db.Scopes(info.Search())
	err = db.Count(&total).Scopes(common.Paginate(&info.PageInfo)).Find(&entities).Error
	return entities, total, err
}
