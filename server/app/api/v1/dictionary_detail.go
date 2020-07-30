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

// CreateDictionaryDetail Create DictionaryDetail
// CreateDictionaryDetail 创建DictionaryDetail
func CreateDictionaryDetail(r *ghttp.Request) {
	var createInfo request.CreateDictionaryDetail
	if err := r.Parse(&createInfo); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.CreateDictionaryDetail(&createInfo); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "创建成功")
}

// DeleteDictionaryDetail Delete DictionaryDetail
// DeleteDictionaryDetail 删除DictionaryDetail
func DeleteDictionaryDetail(r *ghttp.Request) {
	var delete request.DeleteById
	if err := r.Parse(&delete); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.DeleteDictionaryDetail(&delete); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// UpdateDictionaryDetail Update DictionaryDetail
// UpdateDictionaryDetail 更新DictionaryDetail
func UpdateDictionaryDetail(r *ghttp.Request) {
	var update request.UpdateDictionaryDetail
	if err := r.Parse(&update); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.UpdateDictionaryDetail(&update); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("更新失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "更新成功")
}

// FindDictionaryDetail Query DictionaryDetail with id
// FindDictionaryDetail 用id查询DictionaryDetail
func FindDictionaryDetail(r *ghttp.Request) {
	var find request.FindById
	if err := r.Parse(&find); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	dataReturn, err := service.FindDictionaryDetail(&find)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("查询失败，err:%v", err))
		r.Exit()
	}
	global.OkDetailed(r, g.Map{"resysDictionaryDetail": dataReturn}, "查询成功")
}

// GetDictionaryDetailList Paging to get a list of DictionaryDetails
// GetDictionaryDetailList 分页获取DictionaryDetail列表
func GetDictionaryDetailList(r *ghttp.Request) {
	var pageInfo request.GetDictionaryDetailList
	if err := r.Parse(&pageInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	condition := g.Map{}
	if pageInfo.Status == true || pageInfo.Status == false {
		condition["status"] = utils.BoolToInt(pageInfo.Status)
	}
	if r.GetString("status") == "empty" || r.GetString("status") == "" {
		delete(condition, "status")
	}
	list, total, err := service.GetDictionaryDetailList(&pageInfo, condition)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，err:%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize})
}
