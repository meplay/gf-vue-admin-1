package service

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
)

func (a *api) A1Data() []string {
	return []string{"ID", "api路径", "请求方法", "api组", "api中文描述"}
}

func (a *api) FilePath() string {
	return "./public/excel/" + a._api.TableName() + ".xlsx"
}

func (a *api) DataList() [][]interface{} {
	var apis []model.Api
	switch global.Config.System.DbType {
	case "gdb":
		if err := g.DB().Table(a._api.TableName()).Structs(&apis); err != nil {
			return [][]interface{}{}
		}
	case "gorm":
		if err := global.Db.Find(&apis).Error; err != nil {
			return [][]interface{}{}
		}
	default:
		return [][]interface{}{}
	}

	list2 := make([][]interface{}, 0, len(apis))
	for _, m := range apis {
		list1 := make([]interface{}, 0)
		list1 = append(list1, m.ID)
		list1 = append(list1, m.Path)
		list1 = append(list1, m.Method)
		list1 = append(list1, m.ApiGroup)
		list1 = append(list1, m.Description)
		list2 = append(list2, list1)
	}
	return list2
}

func (a *api) SheetName() string {
	return a._api.TableName()
}
