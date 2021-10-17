package example

import (
    "github.com/flipped-aurora/gf-vue-admin/app/model/example"
    "github.com/flipped-aurora/gf-vue-admin/app/model/example/request"
    "github.com/flipped-aurora/gf-vue-admin/library/common"
    "github.com/flipped-aurora/gf-vue-admin/library/global"
)

var {{.StructName}} = new({{.Abbreviation}})

type {{.Abbreviation}} struct{}

// Create 创建{{.Description}}记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *{{.Abbreviation}}) Create(info *request.{{.StructName}}Create) error {
	return global.Db.Create(&info.{{.StructName}}).Error
}

// First 根据id获取{{.Description}}记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *{{.Abbreviation}}) First(info *common.GetByID) (data *example.{{.StructName}}, err error) {
    var entity example.{{.StructName}}
    err = global.Db.Where("id = ?", info.ID).First(&entity).Error
    return &entity, err
}

// Update 更新{{.Description}}记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *{{.Abbreviation}}) Update(info *request.{{.StructName}}Update) error {
    return global.Db.Where("id = ?", info.ID).Updates(&info.{{.StructName}}).Error
}

// Delete 删除{{.Description}}记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *{{.Abbreviation}}) Delete(info *common.GetByID) (err error) {
	return global.Db.Delete(&example.{{.StructName}}{}, info.ID).Error
}

// Deletes 批量删除{{.Description}}记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *{{.Abbreviation}}) Deletes(ids *common.GetByIDs) error {
	return global.Db.Delete(&[]example.{{.StructName}}{},"id in ?",ids.Ids).Error
}

// GetList 分页获取{{.Description}}记录
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *{{.Abbreviation}}) GetList(info *request.{{.StructName}}Search) (list []example.{{.StructName}}, total int64, err error) {
    entities := make([]example.{{.StructName}}, 0, info.PageSize)
    db := global.Db.Model(&example.{{.StructName}}{})
    db = db.Scopes(info.Search())
	err = db.Count(&total).Find(&entities).Error
	return entities, total, err
}
