package request

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/example"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
    "gorm.io/gorm"
)

type {{.StructName}}Create struct {
    example.{{.StructName}}
}

type {{.StructName}}Update struct {
    example.{{.StructName}}
}

type {{.StructName}}Search struct{
    example.{{.StructName}}
common.PageInfo
}

func (r *{{.StructName}}Search) Search() func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB { // 如果有条件搜索 下方会自动创建搜索语句
    {{- range .Fields}}
        {{- if .FieldSearchType}}
            {{- if eq .FieldType "string" }}
        if info.{{.FieldName}} != "" {
            db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
        }
            {{- else if eq .FieldType "bool" }}
        if info.{{.FieldName}} != nil {
            db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
        }
            {{- else if eq .FieldType "int" }}
        if info.{{.FieldName}} != nil {
            db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
        }
            {{- else if eq .FieldType "float64" }}
        if info.{{.FieldName}} != nil {
            db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
        }
            {{- else if eq .FieldType "time.Time" }}
        if info.{{.FieldName}} != nil {
            db = db.Where("`{{.ColumnName}}` {{.FieldSearchType}} ?",{{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }})
        }
            {{- end }}
        {{- end }}
    {{- end }}
        return db
    }
}