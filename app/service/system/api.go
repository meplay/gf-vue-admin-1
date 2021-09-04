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

var Api = new(api)

type api struct{}

// Create 新增基础api
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) Create(info *request.ApiCreate) error {
	if !errors.Is(global.Db.Where("path = ? AND method = ?", info.Path, info.Method).First(&system.Api{}).Error, gorm.ErrRecordNotFound) {
		return _errors.New("存在相同api!")
	}
	if err := global.Db.Create(&info).Error; err != nil {
		return _errors.Wrap(err, "创建失败!")
	}
	return nil
}

// First 根据id获取api
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) First(info *common.GetByID) (entity *system.Api, err error) {
	entity = &system.Api{}
	err = global.Db.Where("id = ?", info.ID).First(entity).Error
	return
}

// Update 根据id更新api
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) Update(info *request.ApiUpdate) error {
	var entity system.Api
	if err := global.Db.First(&entity, info.ID).Error; err != nil {
		return _errors.Wrap(err, "找不到记录!")
	}
	if entity.Path != info.Path || entity.Method != info.Method {
		if !errors.Is(global.Db.Where("path = ? AND method = ?", entity.Path, entity.Method).First(&system.Api{}).Error, gorm.ErrRecordNotFound) {
			return _errors.New("存在相同api!")
		}
	}
	err := global.Db.Transaction(func(tx *gorm.DB) error {
		// TODO Casbin.UpdateCasbinApi(entity.Path, info.Path, entity.Method, info.Method)
		entity = info.Update()
		if err := tx.Model(&system.Api{}).Where("id = ?", info.ID).Updates(&info).Error; err != nil {
			return _errors.Wrap(err, "更新api信息失败!")
		}
		return nil
	})
	return err
}

// Delete 删除基础api
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) Delete(info *request.DeleteApi) error {
	if err := global.Db.Delete(&system.Api{}, info.ID).Error; err != nil {
		return _errors.Wrap(err, "删除api失败!")
	}
	// TODO CasbinServiceApp.ClearCasbin(1, api.Path, api.Method)
	return nil
}

// Deletes 批量删除 system.Api
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) Deletes(ids *common.GetByIDs) error {
	return global.Db.Delete(&[]system.Api{}, "id in ?", ids.Ids).Error
}

// GetList 分页获取 []system.Api 数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) GetList(info *request.ApiSearch) (list []system.Api, total int64, err error) {
	entities := make([]system.Api, 0, info.PageSize)
	db := global.Db.Model(&system.Api{})
	db = db.Scopes(info.Search())
	err = db.Count(&total).Scopes(common.Paginate(&info.PageInfo)).Find(&entities).Error
	return entities, total, err
}

// GetAllApis 获取所有的 system.Api
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *api) GetAllApis() (apis []system.Api, err error) {
	err = global.Db.Find(&apis).Error
	return
}
