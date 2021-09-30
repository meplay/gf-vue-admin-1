//go:build mysql
// +build mysql

package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/response"
	"github.com/flipped-aurora/gf-vue-admin/library/constant"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
)

// GetDbs 获取所有数据库名
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) GetDbs() (dbs []response.Db, err error) {
	err = global.Db.Raw(constant.GetDbsSql).Scan(&dbs).Error
	return dbs, err
}

// GetTables 获取指定数据库的所有表名
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) GetTables(dbName string) (tables []response.Table, err error) {
	err = global.Db.Raw(constant.GetTablesSql, dbName).Scan(&tables).Error
	return tables, err
}

// GetColumns 获取指定数据库和指定数据表的所有字段名,类型值等
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) GetColumns(tableName string, dbName string) (columns []response.Column, err error) {
	err = global.Db.Raw(constant.GetColumnsSql, tableName, dbName).Scan(&columns).Error
	return columns, err
}
