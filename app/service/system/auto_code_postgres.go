//+build postgres

package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/response"
	"github.com/flipped-aurora/gf-vue-admin/library/constant"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetDbs 获取所有数据库名
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) GetDbs() (dbs *[]response.Db, err error) {
	dbs = &[]response.Db{}
	err = global.Db.Raw(constant.GetDbsSql).Scan(&dbs).Error
	return dbs, err
}

// GetTables 获取指定数据库的所有表名
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) GetTables(dbName string) (tables *[]response.Table, err error) {
	tables = &[]response.Table{}
	db, _err := gorm.Open(postgres.Open(global.Config.Gorm.Dsn.LinkDsn(global.Config.Gorm.Config, dbName)), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if _err != nil {
		return nil, err
	}
	err = db.Raw(constant.GetTablesSql, dbName, "public").Scan(&tables).Error
	return tables, err
}

// GetColumns 获取指定数据库和指定数据表的所有字段名,类型值等
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCode) GetColumns(tableName string, dbName string) (columns *[]response.Column, err error) {
	columns = &[]response.Column{}
	db, _err := gorm.Open(postgres.Open(global.Config.Gorm.Dsn.LinkDsn(global.Config.Gorm.Config, dbName)), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if _err != nil {
		return nil, err
	}
	err = db.Raw(constant.GetColumnsSql, dbName, "public", tableName).Scan(&columns).Error
	return columns, err
}
