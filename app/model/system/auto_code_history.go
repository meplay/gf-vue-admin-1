package system

import (
	"github.com/flipped-aurora/gf-vue-admin/library/global"
)

type AutoCodeHistory struct {
	global.Model
	Flag          int    `json:"flag" gorm:"column:flag;comment:表示对应状态 0 代表创建, 1 代表回滚 ..."`
	TableName     string `json:"table_name" gorm:"column:table_name;comment:表名"`
	StructName    string `json:"structName" gorm:"column:struct_name;comment:结构体名"`
	StructCNName  string `json:"structCNName" gorm:"column:struct_cn_name;comment:结构体中文名"`
	InjectionMeta string `json:"injectionMeta" gorm:"column:injection_meta;comment:注入的内容 RouterPath@functionName@RouterString;"`

	Apis          AutoCodeApis        `json:"apis" gorm:"type:text;column:apis;comment:api表注册内容"`
	Request       AutoCodeStruct      `json:"requestMeta" gorm:"type:text;column:request;comment:前端传入的结构化信息"`
	Injection     []AutoCodeInjection `json:"injection" gorm:"type:text;column:injection;comment:注入的内容结构体"`
	AutoCodePaths []AutoCodePaths     `json:"autoCodePath" gorm:"type:text;column:auto_code_path;comment:其他meta信息 path;path"`
}
