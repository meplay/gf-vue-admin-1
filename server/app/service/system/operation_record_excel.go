package service

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"github.com/gogf/gf/frame/g"
)

func (r *record) A1Data() []string {
	return []string{"ID", "请求ip", "请求路径", "代理", "请求方法", "请求Body", "响应Body", "错误信息", "请求状态", "用户id", "延迟"}
}

func (r *record) FilePath() string {
	return "./public/excel/" + r._record.TableName() + ".xlsx"

}

func (r *record) DataList() [][]interface{} {
	var records []model.OperationRecord
	switch global.Config.System.DbType {
	case "gdb":
		if err := g.DB().Table(r._record.TableName()).Structs(&records); err != nil {
			return [][]interface{}{}
		}
	case "gorm":
		if err := global.Db.Find(&records).Error; err != nil {
			return [][]interface{}{}
		}
	default:
		return [][]interface{}{}
	}

	list2 := make([][]interface{}, 0, len(records))
	for _, m := range records {
		list1 := make([]interface{}, 0)
		list1 = append(list1, m.ID)
		list1 = append(list1, m.Ip)
		list1 = append(list1, m.Path)
		list1 = append(list1, m.Agent)
		list1 = append(list1, m.Method)
		list1 = append(list1, m.Request)
		list1 = append(list1, m.Response)
		list1 = append(list1, m.ErrorMessage)
		list1 = append(list1, m.Status)
		list1 = append(list1, m.UserID)
		list1 = append(list1, m.Latency)
		list2 = append(list2, list1)
	}
	return list2
}

func (r *record) SheetName() string {
	return r._record.TableName()

}
