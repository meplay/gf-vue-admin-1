package service

import (
	"gf-vue-admin/app/api/request"
	model "gf-vue-admin/app/model/system"
	"github.com/gogf/gf/frame/g"
)

var OperationRecord = new(record)

type record struct {
	_record model.OperationRecord
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建记录
func (r *record) Create(info *request.CreateOperationRecord) error {
	_, err := g.DB().Table(r._record.TableName()).Insert(info.Create())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 根据id获取单条操作记录
func (r *record) First(info *request.GetById) (result *model.OperationRecord, err error) {
	var entity model.OperationRecord
	err = g.DB().Table(r._record.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除操作记录
func (r *record) Delete(info *request.GetById) error {
	_, err := g.DB().Table(r._record.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 批量删除记录
func (r *record) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(r._record.TableName()).Delete(info.Condition())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取操作记录列表
func (r *record) GetList(info *request.SearchOperationRecord) (list interface{}, total int, err error) {
	var records []model.OperationRecord
	db := g.DB().Table(r._record.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&records)
	return records, total, err
}
