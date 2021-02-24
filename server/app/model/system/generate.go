package model

import (
	"errors"
	"text/template"
)

var ErrorAutoMove = errors.New("创建代码成功并移动文件成功")

type AutoCode struct {
	TableName          string  `p:"tableName" json:"tableName"`
	StructName         string  `p:"structName" json:"structName"`
	PackageName        string  `p:"packageName" json:"packageName"`
	Description        string  `p:"description" json:"description"`
	Abbreviation       string  `p:"abbreviation" json:"abbreviation"`
	AutoMoveFile       bool    `p:"autoMoveFile" json:"autoMoveFile"`
	AutoCreateApiToSql bool    `p:"autoCreateApiToSql" json:"autoCreateApiToSql"`
	Fields             []Field `p:"fields" json:"fields"`
}

type Field struct {
	Comment         string `p:"comment" json:"comment"`
	DataType        string `p:"dataType" json:"dataType"`
	DictType        string `p:"dictType" json:"dictType"`
	FieldName       string `p:"fieldName" json:"fieldName"`
	FieldDesc       string `p:"fieldDesc" json:"fieldDesc"`
	FieldType       string `p:"fieldType" json:"fieldType"`
	FieldJson       string `p:"fieldJson" json:"fieldJson"`
	ColumnName      string `p:"columnName" json:"columnName"`
	DataTypeLong    string `p:"dataTypeLong" json:"dataTypeLong"`
	FieldSearchType string `p:"fieldSearchType" json:"fieldSearchType"`
}

type TemplateData struct {
	Template         *template.Template
	LocationPath     string // 模板位置路径
	AutoCodePath     string // 生成代码文件路径
	AutoMoveFilePath string // 移动代码路径
}
