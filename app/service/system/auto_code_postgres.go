//go:build postgres
// +build postgres

package system

import (
	"strings"

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
func (s *autoCode) GetTables(dbName string) (tables []response.Table, err error) {
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
	sql := `
		SELECT columns.COLUMN_NAME                                                                                      as column_name,
		   columns.DATA_TYPE                                                                                        as data_type,
		   CASE
			   columns.DATA_TYPE
			   WHEN 'text' THEN
				   concat_ws('', '', columns.CHARACTER_MAXIMUM_LENGTH)
			   WHEN 'varchar' THEN
				   concat_ws('', '', columns.CHARACTER_MAXIMUM_LENGTH)
			   WHEN 'smallint' THEN
				   concat_ws(',', columns.NUMERIC_PRECISION, columns.NUMERIC_SCALE)
			   WHEN 'decimal' THEN
				   concat_ws(',', columns.NUMERIC_PRECISION, columns.NUMERIC_SCALE)
			   WHEN 'integer' THEN
				   concat_ws('', '', columns.NUMERIC_PRECISION)
			   WHEN 'bigint' THEN
				   concat_ws('', '', columns.NUMERIC_PRECISION)
			   ELSE ''
			   END                                                                                                  AS data_type_long,
		   (select description.description
			from pg_description description
			where description.objoid = (select attribute.attrelid
										from pg_attribute attribute
										where attribute.attrelid =
											  (select oid from pg_class class where class.relname = '@table_name') and attname =columns.COLUMN_NAME )
			  and description.objsubid = (select attribute.attnum
										  from pg_attribute attribute
										  where attribute.attrelid =
												(select oid from pg_class class where class.relname = '@table_name') and attname =columns.COLUMN_NAME )) as column_comment
		FROM INFORMATION_SCHEMA.COLUMNS columns
		WHERE table_catalog = '@table_catalog'
		  and table_schema = 'public'
		  and table_name = '@table_name';
	`
	sql = strings.ReplaceAll(sql, "@table_catalog", dbName)
	sql = strings.ReplaceAll(sql, "@table_name", tableName)
	err = db.Raw(constant.GetColumnsSql, dbName, "public", tableName).Scan(&columns).Error
	return columns, err
}
