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
	FieldName       string `json:"fieldName "json:"fieldName"`
	FieldDesc       string `json:"fieldDesc "json:"fieldDesc"`
	FieldType       string `json:"fieldType "json:"fieldType"`
	FieldJson       string `json:"fieldJson "json:"fieldJson"`
	DataType        string `json:"dataType "json:"dataType"`
	DataTypeLong    string `json:"dataTypeLong "json:"dataTypeLong"`
	Comment         string `json:"comment "json:"comment"`
	ColumnName      string `json:"columnName "json:"columnName"`
	FieldSearchType string `json:"fieldSearchType "json:"fieldSearchType"`
	DictType        string `json:"dictType "json:"dictType"`
}
