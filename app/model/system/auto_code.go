package system

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/utils"
	"github.com/pkg/errors"
	"path/filepath"
)

type AutoCodeStruct struct {
	TableName          string   `json:"tableName"`          // 表名
	StructName         string   `json:"structName"`         // Struct名称
	PackageName        string   `json:"packageName"`        // 文件名称
	Description        string   `json:"description"`        // Struct中文名称
	Abbreviation       string   `json:"abbreviation"`       // Struct简称
	HumpPackageName    string   `json:"humpPackageName"`    // go文件名称
	AutoMoveFile       bool     `json:"autoMoveFile"`       // 是否自动移动文件
	AutoCreateApiToSql bool     `json:"autoCreateApiToSql"` // 是否自动创建api
	Fields             []*Field `json:"fields"`

	Injection []AutoCodeInjection `json:"-" gorm:"-"`
}

type Field struct {
	FieldName       string `json:"fieldName"`       // Field名
	FieldDesc       string `json:"fieldDesc"`       // 中文名
	FieldType       string `json:"fieldType"`       // Field数据类型
	FieldJson       string `json:"fieldJson"`       // FieldJson
	DataType        string `json:"dataType"`        // 数据库字段类型
	DataTypeLong    string `json:"dataTypeLong"`    // 数据库字段长度
	Comment         string `json:"comment"`         // 数据库字段描述
	ColumnName      string `json:"columnName"`      // 数据库字段
	FieldSearchType string `json:"fieldSearchType"` // 搜索条件
	DictType        string `json:"dictType"`        // 字典
}

// Scan 扫描
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *AutoCodeStruct) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		if err := json.Unmarshal(v, a); err != nil {
			return err
		}
	case string:
		if err := json.Unmarshal([]byte(v), a); err != nil {
			return err
		}
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal AutoCodeStruct value:", value))
	}
	return nil
}

// Value 值
// Author [SliverHorn](https://github.com/SliverHorn)
func (a AutoCodeStruct) Value() (driver.Value, error) {
	bytes, err := json.Marshal(&a)
	if err != nil {
		return nil, err
	}
	return driver.Value(bytes), err
}

// TrimSpace 结构体去空格
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *AutoCodeStruct) TrimSpace() {
	if a == nil {
		return
	}
	utils.File.TrimSpace(a)
	for i := 0; i < len(a.Fields); i++ {
		utils.File.TrimSpace(a.Fields[i])
	}
}

// GenerateInjection 生成注入内容
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *AutoCodeStruct) GenerateInjection() []AutoCodeInjection {
	entities := []AutoCodeInjection{
		{
			Filepath:       filepath.Join(global.Config.AutoCode.Root, global.Config.AutoCode.Server.Root, global.Config.AutoCode.Server.Boot, "gorm.go"),
			FunctionName:   "Initialize",
			CodeDataFormat: "new(example.%s),",
		},
		{
			Filepath:       filepath.Join(global.Config.AutoCode.Root, global.Config.AutoCode.Server.Root, global.Config.AutoCode.Server.Boot, "router.go"),
			FunctionName:   "PublicRouter",
			CodeDataFormat: "example.New%sRouter(public).Public().PublicWithoutRecord()",
		},
		{
			Filepath:       filepath.Join(global.Config.AutoCode.Root, global.Config.AutoCode.Server.Root, global.Config.AutoCode.Server.Boot, "router.go"),
			FunctionName:   "PrivateRouter",
			CodeDataFormat: "example.New%sRouter(private).Private().PrivateWithoutRecord()",
		},
	}
	for i := 0; i < 3; i++ {
		entities[i].StructName = a.StructName
		entities[i].CodeData = fmt.Sprintf(entities[i].CodeDataFormat, a.StructName)
	}
	a.Injection = entities
	return entities
}
