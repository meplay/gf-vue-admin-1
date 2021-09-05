package system

import (
	"errors"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"gorm.io/gorm"
)

// AutoCreateApi 自动创建api数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) AutoCreateApi(info *system.AutoCodeStruct) (ids []uint, err error) {
	entities := []system.Api{
		{Path: "/" + info.Abbreviation + "/" + "create", Description: "新增" + info.Description, ApiGroup: info.Abbreviation, Method: "POST"},
		{Path: "/" + info.Abbreviation + "/" + "delete", Description: "删除" + info.Description, ApiGroup: info.Abbreviation, Method: "DELETE"},
		{Path: "/" + info.Abbreviation + "/" + "deletes", Description: "批量删除" + info.Description, ApiGroup: info.Abbreviation, Method: "DELETE"},
		{Path: "/" + info.Abbreviation + "/" + "update", Description: "更新" + info.Description, ApiGroup: info.Abbreviation, Method: "PUT"},
		{Path: "/" + info.Abbreviation + "/" + "first", Description: "根据ID获取" + info.Description, ApiGroup: info.Abbreviation, Method: "GET"},
		{Path: "/" + info.Abbreviation + "/" + "getList", Description: "获取" + info.Description + "列表", ApiGroup: info.Abbreviation, Method: "POST"},
	}
	err = global.Db.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(entities); i++ {
			var entity system.Api
			if errors.Is(tx.Where("path = ? AND method = ?", entities[i].Path, entities[i].Method).First(&entity).Error, gorm.ErrRecordNotFound) {
				if err = tx.Create(&entities[i]).Error; err != nil { // 遇到错误时回滚事务
					return err
				} else {
					ids = append(ids, entities[i].ID)
				}
			}
		}
		return nil
	})
	return ids, err
}
