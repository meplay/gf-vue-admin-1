package request

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"gorm.io/gorm"
)

type AutoCodeHistoryCreate struct {
	system.AutoCodeStruct
	Apis          system.AutoCodeApis
	Injection     system.AutoCodeInjections
	AutoCodePaths system.AutoCodePaths
}

func (r *AutoCodeHistoryCreate) Create() system.AutoCodeHistory {
	entity := system.AutoCodeHistory{
		StructName:    r.AutoCodeStruct.StructName,
		StructCNName:  r.AutoCodeStruct.Description,
		AutoCodePaths: r.AutoCodePaths,
		Apis:          r.Apis,
		Request:       r.AutoCodeStruct,
		Injection:     r.Injection,
	}
	if r.TableName != "" {
		entity.TablesName = r.AutoCodeStruct.TableName
	} else {
		entity.TablesName = r.AutoCodeStruct.StructName
	}
	return entity
}

type AutoCodeHistorySearch struct {
	*common.PageInfo
}

func (s *AutoCodeHistorySearch) Select() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select("id,created_at,updated_at,struct_name,struct_cn_name,flag,table_name")
	}
}

func (s *AutoCodeHistorySearch) Order() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order("updated_at desc")
	}
}
