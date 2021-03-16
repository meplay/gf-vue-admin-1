package service

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
)

func (a *admin) A1Data() []string {
	return []string{"ID", "用户UUID", "用户头像", "用户登录名", "用户登录密码", "用户昵称", "用户角色ID"}
}

func (a *admin) FilePath() string {
	return "./public/excel/" + a._admin.TableName() + ".xlsx"
}

func (a *admin) DataList() [][]interface{} {
	var admins []model.Admin
	switch global.Config.System.DbType {
	case "gdb":
		if err := g.DB().Table(a._admin.TableName()).Structs(&admins); err != nil {
			return [][]interface{}{}
		}
	case "gorm":
		if err := global.Db.Find(&admins).Error; err != nil {
			return [][]interface{}{}
		}
	default:
		return [][]interface{}{}
	}

	list2 := make([][]interface{}, 0, len(admins))
	for _, m := range admins {
		list1 := make([]interface{}, 0)
		list1 = append(list1, m.ID)
		list1 = append(list1, m.Uuid)
		list1 = append(list1, m.Avatar)
		list1 = append(list1, m.Username)
		list1 = append(list1, m.Password)
		list1 = append(list1, m.Nickname)
		list1 = append(list1, m.AuthorityId)
		list2 = append(list2, list1)
	}
	return list2
}

func (a *admin) SheetName() string {
	return a._admin.TableName()
}
