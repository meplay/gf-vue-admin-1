package service

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
)

func (d *detail) A1Data() []string {
	return []string{"ID", "展示值", "字典值", "启用状态", "排序标记"}
}

func (d *detail) FilePath() string {
	return "./public/excel/" + d._detail.TableName() + ".xlsx"
}

func (d *detail) DataList() [][]interface{} {
	var infos []model.DictionaryDetail
	switch global.Config.System.DbType {
	case "gdb":
		if err := g.DB().Table(d._detail.TableName()).Structs(&infos); err != nil {
			return [][]interface{}{}
		}
	case "gorm":
		if err := global.Db.Find(&infos).Error; err != nil {
			return [][]interface{}{}
		}
	default:
		return [][]interface{}{}
	}

	list2 := make([][]interface{}, 0, len(infos))
	for _, m := range infos {
		list1 := make([]interface{}, 0)
		list1 = append(list1, m.ID)
		list1 = append(list1, m.Label)
		list1 = append(list1, m.Status)
		list1 = append(list1, m.Value)
		list1 = append(list1, m.Sort)
		list2 = append(list2, list1)
	}
	return list2
}

func (d *detail) SheetName() string {
	return d._detail.TableName()
}
