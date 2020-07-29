package service

import (
	"fmt"
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
