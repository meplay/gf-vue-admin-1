package service

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
)

func (m *menu) A1Data() []string {
	return []string{"ID", "路由path", "路由name", "父菜单ID", "对应前端文件路径", "排序标记", "是否在列表隐藏", "是否缓存", "是否是基础路由", "菜单图标", "菜单名"}
}

func (m *menu) FilePath() string {
	return "./public/excel/" + m._menu.TableName() + ".xlsx"
}

func (m *menu) DataList() [][]interface{} {
	var menus []model.Menu
	switch global.Config.System.DbType {
	case "gdb":
		if err := g.DB().Table(m._menu.TableName()).Structs(&menus); err != nil {
			return [][]interface{}{}
		}
	case "gorm":
		if err := global.Db.Find(&menus).Error; err != nil {
			return [][]interface{}{}
		}
	default:
		return [][]interface{}{}
	}

	list2 := make([][]interface{}, 0, len(menus))
	for _, data := range menus {
		list1 := make([]interface{}, 0)
		list1 = append(list1, data.ID)
		list1 = append(list1, data.Path)
		list1 = append(list1, data.Name)
		list1 = append(list1, data.ParentId)
		list1 = append(list1, data.Component)
		list1 = append(list1, data.Sort)
		list1 = append(list1, data.Hidden)
		list1 = append(list1, data.KeepAlive)
		list1 = append(list1, data.DefaultMenu)
		list1 = append(list1, data.Icon)
		list1 = append(list1, data.Title)
		list2 = append(list2, list1)
	}
	return list2
}

func (m *menu) SheetName() string {
	return m._menu.TableName()
}
