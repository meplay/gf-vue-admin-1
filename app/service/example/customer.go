package example

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/example"
	"github.com/flipped-aurora/gf-vue-admin/app/model/example/request"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
)

var Customer = new(customer)

type  customer struct {}

// First 根据id获取客户表记录
// Author [SliverHorn](https://github.com/Cbzy)
func (e *customer) First(info *common.GetByID) (data *example.Customer,err error){
	var code example.Customer
	err = global.Db.Where("id=?",info.ID).First(&code).Error
	return &code, err
}

// Create 创建客户表记录
// Author [SliverHorn](https://github.com/Cbzy)
func (e *customer) Create(info *request.CustomerCreate) error {
	return global.Db.Create(&info.Customer).Error
}

// Update 更新客户表记录
// Author [SliverHorn](https://github.com/Cbzy)
func (e *customer) Update(info *request.CustomerUpdate) error{
	return global.Db.Where("id = ?" ,&info.ID).Updates(&info.Customer).Error
}

// Delete 删除客户表记录
// Author [SliverHorn](https://github.com/Cbzy)
func (e *customer) Delete(info *common.GetByID) (err error) {
 return global.Db.Delete(&example.Customer{},info.ID).Error
}

// Deletes 批量删除客户表记录
// Author [SliverHorn](https://github.com/Cbzy)
func (e *customer) Deletes(ids *common.GetByIDs) error{
	return global.Db.Delete(&[]example.Customer{},"id in ?",ids.Ids).Error
}

// GetList 分页获取客户表记录
// Author [SliverHorn](https://github.com/Cbzy)
func (e *customer) GetList(info *request.CustomerSearch) (list []example.Customer,total int64,err error) {
	entities := make([]example.Customer,0,info.PageSize)
	db := global.Db.Model(&example.Customer{})
	db = db.Scopes(info.Search())
	err = db.Count(&total).Find(&entities).Error
	return entities,total,err
}