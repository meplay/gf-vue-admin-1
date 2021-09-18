package system

import "github.com/flipped-aurora/gf-vue-admin/library/utils"

type AutoCodeStruct struct {
	StructName         string   `json:"structName"`         // Struct名称
	TableName          string   `json:"tableName"`          // 表名
	PackageName        string   `json:"packageName"`        // 文件名称
	HumpPackageName    string   `json:"humpPackageName"`    // go文件名称
	Abbreviation       string   `json:"abbreviation"`       // Struct简称
	Description        string   `json:"description"`        // Struct中文名称
	AutoCreateApiToSql bool     `json:"autoCreateApiToSql"` // 是否自动创建api
	AutoMoveFile       bool     `json:"autoMoveFile"`       // 是否自动移动文件
	Fields             []*Field `json:"fields"`
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
