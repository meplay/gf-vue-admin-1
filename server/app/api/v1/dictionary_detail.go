package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/service"
	"server/library/global"

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
	var deleteRequest request.DeleteDictionaryDetail
	if err := r.Parse(&deleteRequest); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.DeleteDictionaryDetail(&deleteRequest); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// UpdateDictionaryDetail Update DictionaryDetail
// UpdateDictionaryDetail 更新DictionaryDetail
func UpdateDictionaryDetail(r *ghttp.Request) {
	var updateRequest request.UpdateDictionaryDetail
	if err := r.Parse(&updateRequest); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	if err := service.UpdateDictionaryDetail(&updateRequest); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("更新失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "更新成功")
}

// FindDictionaryDetail Query DictionaryDetail with id
// FindDictionaryDetail 用id查询DictionaryDetail
func FindDictionaryDetail(r *ghttp.Request) {
	var findRequest request.FindDictionaryDetail
	if err := r.Parse(&findRequest); err != nil {
		global.FailWithMessage(r, err.Error())
	}
	dataReturn, err := service.FindDictionaryDetail(&findRequest)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("查询失败，err:%v", err))
		r.Exit()
	}
	global.OkDetailed(r, g.Map{"DictionaryDetail": dataReturn}, "查询成功")
}

// GetDictionaryDetailList Paging to get a list of DictionaryDetails
// GetDictionaryDetailList 分页获取DictionaryDetail列表
func GetDictionaryDetailList(r *ghttp.Request) {
	var pageInfoList request.GetDictionaryDetailList
	if err := r.Parse(&pageInfoList); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetDictionaryDetailList(&pageInfoList)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，err:%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.PageResult{List: list, Total: total, Page: pageInfoList.Page, PageSize: pageInfoList.PageSize})
}
