package request

import (
	"gf-vue-admin/app/model"
	"github.com/gogf/gf/frame/g"
)

type Search{{.StructName}} struct{
    model.{{.StructName}}
    PageInfo
}

func (s *Search{{.StructName}}) Search() g.Map {
	condition := g.Map{}
	{{- range .Fields}}
                {{- if .FieldSearchType}}
                    {{- if eq .FieldType "string" }}
        if s.{{.FieldName}} != "" {
            condition["`{{.ColumnName}}`{{if eq .FieldSearchType "LIKE"}} {{.FieldSearchType}} ?{{ end }}"] = {{if eq .FieldSearchType "LIKE"}}"%"+{{ end }} + s.{{.FieldName}} + {{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}
        }
                    {{- else if eq .FieldType "bool" }}
        if info.{{.FieldName}} != nil {
            if *s.{{.FieldName}} == true {
                condition["`{{.FieldName}}`"] = 1
            } else {
                condition["`{{.FieldName}}`"] = 2
            }
        }
                    {{- else if eq .FieldType "int" }}
        if info.{{.FieldName}} != 0 {
            condition["`{{.ColumnName}}` {{.FieldSearchType}} ?"] = {{if eq .FieldSearchType "LIKE"}}"%"+{{ end }} + s.{{.FieldName}} + {{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}
        }
                    {{- else if eq .FieldType "float64" }}
        if info.{{.FieldName}} != 0 {
            condition["`{{.ColumnName}}`{{if eq .FieldSearchType "LIKE"}} {{.FieldSearchType}} ?{{ end }}"] = {{if eq .FieldSearchType "LIKE"}}"%"+{{ end }} + s.{{.FieldName}} + {{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}
        }
                    {{- else if eq .FieldType "time.Time" }}
        if !info.{{.FieldName}}.IsZero() {
            condition["`{{.ColumnName}}`{{if eq .FieldSearchType "LIKE"}} {{.FieldSearchType}} ?{{ end }}"] = {{if eq .FieldSearchType "LIKE"}}"%"+{{ end }} + s.{{.FieldName}} + {{if eq .FieldSearchType "LIKE"}}"%"+{{ end }}
        }
                    {{- end }}
            {{- end }}
        {{- end }}
    return condition
}