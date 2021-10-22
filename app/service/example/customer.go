package example

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/example"
	"github.com/flipped-aurora/gf-vue-admin/app/model/example/request"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
)

var Customer = new(customer)

type customer struct{}

// Create 创建客户表记录
// Author [SliverHorn](https://github.com/Cbzy)
func (e *customer) Create(info *request.CustomerCreate) error {
	return global.Db.Create(&info.Customer).Error
}

// First 根据id获取客户表记录
// Author [SliverHorn](https://github.com/Cbzy)
func (e *customer) First(info *common.GetByID) (data *example.Customer, err error) {
	var entity example.Customer
	err = global.Db.Where("id = ?", info.ID).First(&entity).Error
	return &entity, err
}

// Update 更新客户表记录
// Author [SliverHorn](https://github.com/Cbzy)
func (e *customer) Update(info *request.CustomerUpdate) error {
	return global.Db.Where("id = ?", info.ID).Updates(&info.Customer).Error
}

// Delete 删除客户表记录
// Author [SliverHorn](https://github.com/Cbzy)
func (e *customer) Delete(info *common.GetByID) error {
	return global.Db.Delete(&example.Customer{}, info.ID).Error
}

// GetList 分页获取客户表记录
// Author [SliverHorn](https://github.com/Cbzy)
func (e *customer) GetList(info *request.CustomerSearch) (list []example.Customer, total int64, err error) {
	var ids []string
	if err = global.Db.Model(&system.AuthoritiesResources{}).Select("resources_id").Where("authority_id = ?", info.UserAuthorityID).Scan(&ids).Error; err != nil {
		return nil, total, errors.Wrap(err, "查找资源权限ids失败")
	}
	entities := make([]example.Customer, 0, info.PageSize)
	db := global.Db.Model(&example.Customer{})
	db = db.Where("user_authority_id in ?", ids).Scopes(info.Search())
	err = db.Count(&total).Find(&entities).Error
	return entities, total, err
}
