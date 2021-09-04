package system

import (
	"errors"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	_errors "github.com/pkg/errors"
	"gorm.io/gorm"
)

var Dictionary = new(dictionary)

type dictionary struct{}

// Create 创建字典数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionary) Create(info *request.DictionaryCreate) error {
	if (!errors.Is(global.Db.First(&system.Dictionary{}, "type = ?", info.Type).Error, gorm.ErrRecordNotFound)) {
		return _errors.New("存在相同的type，不允许创建!")
	}
	entity := info.Create()
	return global.Db.Create(&entity).Error
}

// First 根据id或者type获取字典单条数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionary) First(info *request.DictionaryFirst) (data *system.Dictionary, err error) {
	var entity system.Dictionary
	err = global.Db.Where("type = ? OR id = ?", info.Type, info.ID).Preload("DictionaryDetails").First(&entity).Error
	return &entity, err
}

// Update 更新字典数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionary) Update(info *request.DictionaryUpdate) error {
	var entity system.Dictionary
	update := info.Update()
	if err := global.Db.Where("id = ?", info.ID).First(&entity).Error; err != nil {
		return _errors.New("找不到记录!")
	}
	if entity.Type != info.Type {
		if !errors.Is(global.Db.First(&system.Dictionary{}, "type = ?", info.Type).Error, gorm.ErrRecordNotFound) {
			return _errors.New("存在相同的type，不允许创建!")
		}
	}
	if err := global.Db.Model(&system.Dictionary{}).Updates(update).Error; err != nil {
		return _errors.Wrap(err, "更新失败!")
	}
	return nil
}

// Delete 删除字典数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionary) Delete(info *common.GetByID) error {
	var entity system.Dictionary
	if err := global.Db.First(&entity, info.ID).Error; err != nil {
		return _errors.Wrap(err, "非法删除!")
	}
	if err := global.Db.Delete(&entity).Delete(&entity.DictionaryDetails).Error; err != nil {
		return _errors.Wrap(err, "删除失败!")
	}
	return nil
}

// GetList 分页获取字典列表
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *dictionary) GetList(info *request.DictionarySearch) (list []system.Dictionary, total int64, err error) {
	entities := make([]system.Dictionary, 0, info.PageSize)
	db := global.Db.Model(&system.Dictionary{})
	err = db.Count(&total).Scopes(info.Search()).Find(&entities).Error
	return entities, total, err
}
