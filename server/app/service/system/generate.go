package service

import (
	"errors"
	"gf-vue-admin/app/api/response"
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/app/service/system/internal"
	"gf-vue-admin/library/constant"
	"gf-vue-admin/library/utils"
	"github.com/gogf/gf/frame/g"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var Generate = new(generate)

type generate struct {
	err  error
	file *os.File
	data []model.TemplateData // 模板数据切片

	content   []byte   // 生成数据
	files     []string // 全部模板文件名
	directory []string // 需要创建的文件夹/目录
	extension string   // 文件后缀名
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取所有的数据库名
func (s *generate) GetDB() (result []*response.Dbs, err error) {
	err = g.DB().GetStructs(&result, "SELECT SCHEMA_NAME AS `database` FROM INFORMATION_SCHEMA.SCHEMATA;")
	return result, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取数据库的所有表名
func (s *generate) GetTables(db string) (result []response.Tables, err error) {
	err = g.DB().GetStructs(&result, "select table_name as table_name from information_schema.tables where table_schema = ?", db)
	return result, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取指定数据库与指定表名的表字段
func (s *generate) GetColumn(db string, table string) (result []*response.Columns, err error) {
	err = g.DB().GetStructs(&result, "SELECT COLUMN_NAME column_name,DATA_TYPE data_type,CASE DATA_TYPE WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH WHEN 'varchar' THEN c.CHARACTER_MAXIMUM_LENGTH WHEN 'double' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'decimal' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'int' THEN c.NUMERIC_PRECISION WHEN 'bigint' THEN c.NUMERIC_PRECISION ELSE '' END AS data_type_long,COLUMN_COMMENT colume_comment FROM INFORMATION_SCHEMA.COLUMNS c WHERE table_name = ? AND table_schema = ?", table, db)
	return result, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建代码
func (s *generate) Create(info *model.AutoCode) error {

	if s.data, s.files, s.directory, s.err = internal.Generate.GetNeedList(info); s.err != nil { // 获取需要文件内容
		return s.err
	}

	if s.err = utils.Directory.BatchCreate(s.directory...); s.err != nil { // 写入文件前，先创建文件夹
		return s.err
	}

	// 生成文件
	for _, value := range s.data {
		if s.file, s.err = os.OpenFile(value.AutoCodePath, os.O_CREATE|os.O_WRONLY, 0755); s.err != nil {
			return s.err
		}
		if s.err = value.Template.Execute(s.file, info); s.err != nil {
			return s.err
		}
		_ = s.file.Close()
	}

	defer func() { // 移除中间文件
		_ = os.RemoveAll(constant.AutoPath)
	}()

	if info.AutoMoveFile { // 判断是否需要自动转移

		for index := range s.data {
			internal.Generate.AddAutoMoveFile(&s.data[index])
		}

		for _, value := range s.data { // 移动文件
			if err := internal.Generate.FileMove(value.AutoCodePath, value.AutoMoveFilePath); err != nil {
				return err
			}
		}

		return errors.New("创建代码成功并移动文件成功")
	} else { // 打包

		if err := utils.ZipFiles("./generate.zip", s.files, ".", "."); err != nil {
			return err
		}

		return nil
	}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 预览创建代码
func (s *generate) Preview(info *model.AutoCode) (result map[string]string, err error) {

	result = make(map[string]string, len(s.data)) // 初始化map

	if s.data, _, s.directory, s.err = internal.Generate.GetNeedList(info); s.err != nil { // 获取需要文件内容
		return result, s.err
	}

	if err = utils.Directory.BatchCreate(s.directory...); err != nil { // 写入文件前，先创建文件夹
		return result, err
	}

	for _, value := range s.data {
		if s.extension = filepath.Ext(value.AutoCodePath); s.extension == ".txt" {
			continue
		}

		if s.file, s.err = os.OpenFile(value.AutoCodePath, os.O_CREATE|os.O_WRONLY, 0755); s.err != nil {
			return result, err
		}

		if s.err = value.Template.Execute(s.file, info); s.err != nil {
			return result, err
		}

		_ = s.file.Close()

		if s.file, s.err = os.OpenFile(value.AutoCodePath, os.O_CREATE|os.O_RDONLY, 0755); s.err != nil {
			return result, err
		}

		builder := strings.Builder{}
		builder.WriteString("```")

		if s.extension != "" && strings.Contains(s.extension, ".") {
			builder.WriteString(strings.Replace(s.extension, ".", "", -1))
		}

		builder.WriteString("\n\n")

		if s.content, s.err = ioutil.ReadAll(s.file); s.err != nil {
			return result, err
		}

		builder.Write(s.content)
		builder.WriteString("\n\n```")
		pathArr := strings.Split(value.AutoCodePath, string(os.PathSeparator))
		result[pathArr[1]+"-"+pathArr[3]] = builder.String()
		_ = s.file.Close()
	}

	defer func() { // 移除中间文件
		_ = os.RemoveAll(constant.AutoPath)
	}()

	return result, nil
}
