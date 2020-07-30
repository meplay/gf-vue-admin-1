package service

import (
	"fmt"
	"server/app/model"
	"testing"
)

func TestGetDB(t *testing.T) {
	dbs, err := GetDB()
	if err != nil {
		panic(err)
	}
	for _, db := range dbs {
		fmt.Println(db.Database)
	}
}

func TestGetTables(t *testing.T) {
	tables, err := GetTables("gf-vue-admin")
	if err != nil {
		panic(err)
	}
	for _, table := range tables {
		fmt.Println(table.TableName)
	}
}

func TestGetColumns(t *testing.T) {
	columns, err := GetColumn("gf-vue-admin", "admins")
	if err != nil {
		panic(err)
	}
	for _, column := range columns {
		fmt.Println(column.ColumnName)
		fmt.Println(column.DataType)
		fmt.Println(column.DataTypeLong)
		fmt.Println(column.ColumnComment)
	}
}

func TestCreateTemp(t *testing.T) {
	err := CreateTemp(model.AutoCodeStruct{
		StructName:         "CasbinRule",
		TableName:          "casbin_rule",
		PackageName:        "casbin_rule",
		Abbreviation:       "casbinRule",
		Description:        "casbinRule",
		AutoCreateApiToSql: true,
		Fields: []model.Field{
			{
				FieldName:       "PType",
				FieldDesc:       "pType字段",
				FieldType:       "string",
				DataType:        "varchar",
				FieldJson:       "pType",
				DataTypeLong:    "100",
				ColumnName:      "p_type",
				Comment:         "pType字段",
				FieldSearchType: "",
				DictType:        "",
			},
			{
				FieldName:       "V0",
				FieldDesc:       "v0字段",
				FieldType:       "string",
				DataType:        "varchar",
				FieldJson:       "v0",
				DataTypeLong:    "100",
				ColumnName:      "v0",
				Comment:         "v0字段",
				FieldSearchType: "",
				DictType:        "",
			},
			{
				FieldName:       "V1",
				FieldDesc:       "v1字段",
				FieldType:       "string",
				DataType:        "varchar",
				FieldJson:       "v1",
				DataTypeLong:    "100",
				ColumnName:      "v1",
				Comment:         "v1字段",
				FieldSearchType: "",
				DictType:        "",
			},
			{
				FieldName:       "V2",
				FieldDesc:       "v2字段",
				FieldType:       "string",
				DataType:        "varchar",
				FieldJson:       "v2",
				DataTypeLong:    "100",
				ColumnName:      "v2",
				Comment:         "v2字段",
				FieldSearchType: "",
				DictType:        "",
			},
			{
				FieldName:       "V3",
				FieldDesc:       "v3字段",
				FieldType:       "string",
				DataType:        "varchar",
				FieldJson:       "v3",
				DataTypeLong:    "100",
				ColumnName:      "v3",
				Comment:         "v3字段",
				FieldSearchType: "",
				DictType:        "",
			},
			{
				FieldName:       "V4",
				FieldDesc:       "v4字段",
				FieldType:       "string",
				DataType:        "varchar",
				FieldJson:       "v4",
				DataTypeLong:    "100",
				ColumnName:      "v4",
				Comment:         "v4字段",
				FieldSearchType: "",
				DictType:        "",
			},
			{
				FieldName:       "V5",
				FieldDesc:       "v5字段",
				FieldType:       "string",
				DataType:        "varchar",
				FieldJson:       "v5",
				DataTypeLong:    "100",
				ColumnName:      "v5",
				Comment:         "v5字段",
				FieldSearchType: "",
				DictType:        "",
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
