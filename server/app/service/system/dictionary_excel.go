package service

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
)

func (d *dictionary) A1Data() []string {
	return []string{"ID", "字典名(中)", "字典名(英)", "状态", "描述"}
}

func (d *dictionary) FilePath() string {
	return "./public/excel/" + d._dictionary.TableName() + ".xlsx"
}

func (d *dictionary) DataList() [][]interface{} {
	var dictionaries []model.Dictionary
	switch global.Config.System.DbType {
	case "gdb":
		if err := g.DB().Table(d._dictionary.TableName()).Structs(&dictionaries); err != nil {
			return [][]interface{}{}
		}
	case "gorm":
		if err := global.Db.Find(&dictionaries).Error; err != nil {
			return [][]interface{}{}
		}
	default:
		return [][]interface{}{}
	}

	list2 := make([][]interface{}, 0, len(dictionaries))
	for _, m := range dictionaries {
		list1 := make([]interface{}, 0)
		list1 = append(list1, m.ID)
		list1 = append(list1, m.Name)
		list1 = append(list1, m.Type)
		list1 = append(list1, m.Status)
		list1 = append(list1, m.Desc)
		list2 = append(list2, list1)
	}
	return list2
}

func (d *dictionary) SheetName() string {
	return d._dictionary.TableName()
}
