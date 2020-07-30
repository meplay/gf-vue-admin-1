package service

import (
	"server/app/api/request"
	"server/app/model/admins"
	"server/app/model/{{.TableName}}"
	"server/library/global"

	"github.com/gogf/gf/frame/g"
)

// Create{{.StructName}} create a {{.StructName}}
// Create{{.StructName}} 创建{{.StructName}}
func Create{{.StructName}}(create *request.Create{{.StructName}}) (err error) {
	insert := {{.TableName}}.Entity{
	{{- range .Fields}}
        {{- if ne .FieldName "Id" "ID" "CreateAt" "UpdateAt" "DeleteAt"}}
        {{.ColumnName}}: create.{{.FieldName}},
        {{- end}}
    {{- end }}
	}
	_, err = operations.Insert(&insert)
	return err
}

// Delete{{.StructName}} delete {{.StructName}}
// Delete{{.StructName}} 删除 {{.StructName}}
func Delete{{.StructName}}(delete *request.DeleteById) (err error) {
	_, err = operations.Delete(g.Map{"id": delete.Id})
	return err
}

// Delete{{.StructName}}s batch deletion {{.StructName}}s
// Delete{{.StructName}}s 批量删除 {{.StructName}}s
func Delete{{.StructName}}s(deletes *request.DeleteByIds) (err error) {
	_, err = operations.Delete(g.Map{"id IN(?)": deletes.Ids})
	return err
}

// Update{{.StructName}} update {{.StructName}}s
// Update{{.StructName}} 更新 {{.StructName}}s
func Update{{.StructName}}(update *request.Update{{.StructName}}) (err error) {
	condition := g.Map{"id": update.Id}
	updateData := g.Map{
    {{- range .Fields}}
        {{- if ne .FieldName "Id" "ID" "CreateAt" "UpdateAt" "DeleteAt"}}
        {{.ColumnName}}: update.{{.FieldName}},
        {{- end}}
    {{- end }}
	}
	_, err = operations.Update(updateData, condition)
	return err
}

// Find{{.StructName}} Gets a single {{.StructName}} based on id
// Find{{.StructName}} 根据id获取单条{{.StructName}}
func Find{{.StructName}}(find *request.Find{{.StructName}}) (data *{{.TableName}}.Entity, err error) {
	return operations.FindOne(g.Map{"id": find.Id})
}

// Get{{.StructName}}List Page out the {{.StructName}} list
// Get{{.StructName}}List 分页获取{{.StructName}}列表
func Get{{.StructName}}List(info *request.Get{{.StructName}}List, condition g.Map) (list interface{}, total int, err error) {
	datalist := ([]*{{.TableName}}.{{.StructName}})(nil)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.DB(global.Db).Table("{{.TableName}}").Safe()
	{{- range .Fields}}
    	    {{- if .FieldSearchType}}
    	        {{- if eq .FieldType "string" }}
    if pageInfo.{{.ColumnName}} != "" {
        condition["`{{.ColumnName}}` {{.FieldSearchType}} ?"] = {{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }}
    }
    	        {{- else if eq .FieldType "int" }}
    if pageInfo.{{.ColumnName}} != 0 {
        condition["`{{.ColumnName}}` {{.FieldSearchType}} ?"] = {{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }}
    }
                {{- else if eq .FieldType "float64" }}
    if info.{{.FieldName}} != 0 {
        condition["`{{.ColumnName}}` {{.FieldSearchType}} ?"] = {{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }}
    }
                {{- else if eq .FieldType "gtime.Time" }}
    if info.{{.FieldName}}.IsZero() {
        condition["`{{.ColumnName}}` {{.FieldSearchType}} ?"] = {{if eq .FieldSearchType "LIKE"}}"%"+ {{ end }}info.{{.FieldName}}{{if eq .FieldSearchType "LIKE"}}+"%"{{ end }}
    }
            {{- end }}
        {{- end }}
    {{- end }}
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Structs(&datalist)
	return datalist, total, err
}
