package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/service"
	"server/library/global"
	"server/library/utils"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

// CreateDictionary Create Dictionary
// CreateDictionary 创建Dictionary
func CreateDictionary(r *ghttp.Request) {
	var createInfo request.CreateDictionary
	if err := r.Parse(&createInfo); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.CreateDictionary(&createInfo); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "创建成功")
}

// DeleteDictionary Delete Dictionary
// DeleteDictionary 删除Dictionary
func DeleteDictionary(r *ghttp.Request) {
	var deleteInfo request.DeleteDictionary
	if err := r.Parse(&deleteInfo); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.DeleteDictionary(&deleteInfo); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// UpdateDictionary Update Dictionary
// UpdateDictionary 更新Dictionary
func UpdateDictionary(r *ghttp.Request) {
	var updateInfo request.UpdateDictionary
	if err := r.Parse(&updateInfo); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.UpdateDictionary(&updateInfo); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("更新失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "更新成功")
}

// FindDictionary Look up a Dictionary with an ID
// FindDictionary 用id查询Dictionary
func FindDictionary(r *ghttp.Request) {
	var findInfo request.FindDictionary
	if err := r.Parse(&findInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	dictionary, err := service.FindDictionary(&findInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("查询失败，err:%v", err))
		r.Exit()
	}
	global.OkWithData(r, g.Map{"resysDictionary": dictionary})
}

// GetDictionaryList Pagination gets the Dictionary list
// GetDictionaryList 分页获取Dictionary列表
func GetDictionaryList(r *ghttp.Request) {
	var pageInfo request.DictionaryInfoList
	if err := r.Parse(&pageInfo); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	condition := g.Map{}
	if pageInfo.Status == true || pageInfo.Status == false {
		condition["status"] = utils.BoolToInt(pageInfo.Status)
	}
	if r.GetString("status") == "empty" || r.GetString("status") == "" {
		delete(condition, "status")
	}
	list, total, err := service.GetDictionaryInfoList(&pageInfo, condition)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，err:%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize})
}
