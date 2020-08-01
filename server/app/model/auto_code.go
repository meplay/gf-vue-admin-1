package model

// 初始版本自动化代码工具
type AutoCodeStruct struct {
	StructName         string  `p:"structName" json:"structName"`
	TableName          string  `p:"tableName" json:"tableName"`
	PackageName        string  `p:"packageName" json:"packageName"`
	Abbreviation       string  `p:"abbreviation" json:"abbreviation"`
	Description        string  `p:"description" json:"description"`
	AutoCreateApiToSql bool    `p:"autoCreateApiToSql" json:"autoCreateApiToSql"`
	Fields             []Field `p:"fields" json:"fields"`
}

type Field struct {
	FieldName       string `p:"fieldName " json:"fieldName"`
	FieldDesc       string `p:"fieldDesc " json:"fieldDesc"`
	FieldType       string `p:"fieldType " json:"fieldType"`
	FieldJson       string `p:"fieldJson " json:"fieldJson"`
	DataType        string `p:"dataType " json:"dataType"`
	DataTypeLong    string `p:"dataTypeLong " json:"dataTypeLong"`
	Comment         string `p:"comment " json:"comment"`
	ColumnName      string `p:"columnName " json:"columnName"`
	FieldSearchType string `p:"fieldSearchType " json:"fieldSearchType"`
	DictType        string `p:"dictType " json:"dictType"`
}
