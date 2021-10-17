package example

import "github.com/flipped-aurora/gf-vue-admin/library/global"

// {{.StructName}} 结构体
type {{.StructName}} struct {
      global.Model
      {{- range .Fields}}
            {{- if ne .FieldType "string" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- end }}"`
            {{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if ne .FieldType "string" -}}{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}{{- end -}}"`
            {{- end }} {{- end }}
}

{{ if .TableName }}
// TableName {{.StructName}} 表名
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}
