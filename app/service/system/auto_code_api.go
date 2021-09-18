package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
)

// AutoCreateApi 自动创建api数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) AutoCreateApi(info *system.AutoCodeStruct) (entities []system.Api, err error) {
	entities = []system.Api{
		{Path: "/" + info.Abbreviation + "/" + "create", Description: "新增" + info.Description, ApiGroup: info.Abbreviation, Method: "POST"},
		{Path: "/" + info.Abbreviation + "/" + "first", Description: "根据ID获取" + info.Description, ApiGroup: info.Abbreviation, Method: "GET"},
		{Path: "/" + info.Abbreviation + "/" + "update", Description: "更新" + info.Description, ApiGroup: info.Abbreviation, Method: "PUT"},
		{Path: "/" + info.Abbreviation + "/" + "delete", Description: "删除" + info.Description, ApiGroup: info.Abbreviation, Method: "DELETE"},
		{Path: "/" + info.Abbreviation + "/" + "deletes", Description: "批量删除" + info.Description, ApiGroup: info.Abbreviation, Method: "DELETE"},
		{Path: "/" + info.Abbreviation + "/" + "getList", Description: "获取" + info.Description + "列表", ApiGroup: info.Abbreviation, Method: "POST"},
	}
	err = global.Db.Create(&entities).Error
	return entities, err
}
