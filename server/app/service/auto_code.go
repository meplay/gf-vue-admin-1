package service

import (
	"io/ioutil"
	"server/app/api/request"
	"server/app/model"
	"strings"

	"github.com/gogf/gf/frame/g"
)

func CreateTemp(autoCode model.AutoCodeStruct) (err error) {
	//basePath := "./public/template"
	return
}

// GetAllTplFile 用来获取 pathName 文件夹下所有 tpl 文件
func GetAllTplFile(pathName string, fileList []string) ([]string, error) {
	files, err := ioutil.ReadDir(pathName)
	for _, fi := range files {
		if fi.IsDir() {
			fileList, err = GetAllTplFile(pathName+"/"+fi.Name(), fileList)
			if err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(fi.Name(), ".tpl") {
				fileList = append(fileList, pathName+"/"+fi.Name())
			}
		}
	}
	return fileList, err
}

// GetTables Get all Tables of the specified database table name
// GetTables 获取指定数据库表名所有的Table
func GetTables(dbName string) (TableNames []*request.TableReq, err error) {
	err = g.DB("default").GetStructs(&TableNames, "select table_name as table_name from information_schema.tables where table_schema = ?", dbName)
	return TableNames, err
}

// GetDB Get all database names
// GetDB 获取所有的数据库名
func GetDB() (DBNames []*request.DBReq, err error) {
	err = g.DB("default").GetStructs(&DBNames, "SELECT SCHEMA_NAME AS `database` FROM INFORMATION_SCHEMA.SCHEMATA;")
	return DBNames, err
}

// GetColumn Get the table fields of the specified database and the specified table name
// GetColumn 获取指定数据库与指定表名的表字段
func GetColumn(dbName string, tableName string) (Columns []request.ColumnReq, err error) {
	err = g.DB("default").GetStructs(&Columns, "SELECT COLUMN_NAME column_name,DATA_TYPE data_type,CASE DATA_TYPE WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH WHEN 'varchar' THEN c.CHARACTER_MAXIMUM_LENGTH WHEN 'double' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'decimal' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'int' THEN c.NUMERIC_PRECISION WHEN 'bigint' THEN c.NUMERIC_PRECISION ELSE '' END AS data_type_long,COLUMN_COMMENT colume_comment FROM INFORMATION_SCHEMA.COLUMNS c WHERE table_name = ? AND table_schema = ?", tableName, dbName)
	return Columns, err
}
