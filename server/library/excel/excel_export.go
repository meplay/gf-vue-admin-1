package excel

import (
	"fmt"
	"gf-vue-admin/library/excel/internal"
	"gorm.io/gorm/schema"
	"reflect"
)

var Export = new(export)

type export struct{}

func (e *export) Export(entity interface{}) {
	_type := reflect.TypeOf(entity)
	a1 := make([]string, 0, _type.NumField())
	for i := 0; i < _type.NumField(); i++ {
		field := _type.Field(i)
		if field.Name == "Model" || field.Name == "global.Model" || field.Name == "gorm.Model" {
			if id, ok := _type.FieldByName("ID"); ok {
				m1 := schema.ParseTagSetting(id.Tag.Get("gorm"), ";")
				a1 = append(a1, internal.Tool.GetA1Value(m1, id.Name))
			}
			if createdAt, ok := _type.FieldByName("CreatedAt"); ok {
				m2 := schema.ParseTagSetting(createdAt.Tag.Get("gorm"), ";")
				a1 = append(a1, internal.Tool.GetA1Value(m2, createdAt.Name))
			}
			if updatedAt, ok := _type.FieldByName("UpdatedAt"); ok {
				m3 := schema.ParseTagSetting(updatedAt.Tag.Get("gorm"), ";")
				a1 = append(a1, internal.Tool.GetA1Value(m3, updatedAt.Name))
			}
			if deletedAt, ok := _type.FieldByName("DeletedAt"); ok {
				m4 := schema.ParseTagSetting(deletedAt.Tag.Get("gorm"), ";")
				a1 = append(a1, internal.Tool.GetA1Value(m4, deletedAt.Name))
			}
			continue
		}
		if field.Type.Kind() == reflect.Struct {
			m5 := schema.ParseTagSetting(field.Tag.Get("gorm"), ";")
			fmt.Println(field.Type.Kind())
			a1 = append(a1, internal.Tool.GetA1Value(m5, field.Name))
			continue
		}
		m5 := schema.ParseTagSetting(field.Tag.Get("gorm"), ";")
		fmt.Println(field.Type.Kind())
		a1 = append(a1, internal.Tool.GetA1Value(m5, field.Name))
	}

	value := reflect.ValueOf(entity) // value

	for i := 0; i < value.NumField(); i++ {
		value.String()
	}
	
	fmt.Println(a1)
}
