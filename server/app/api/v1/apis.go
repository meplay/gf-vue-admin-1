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

// CreateApi Creating the base API
// CreateApi 创建基础api
func CreateApi(r *ghttp.Request) {
	var c request.CreateApi
	if err := r.Parse(&c); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.CreateApi(&c); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "创建成功")
}

// UpdateApi Update the base API
// UpdateApi 更新基础api
func UpdateApi(r *ghttp.Request) {
	var u request.UpdateApi
	if err := r.Parse(&u); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.UpdateApi(&u); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("更新失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "更新成功")
}

// DeleteApi Delete the specified API
// DeleteApi 删除指定api
func DeleteApi(r *ghttp.Request) {
	var d request.DeleteApi
	if err := r.Parse(&d); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteApi(&d); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，err:%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// GetApiById Get the API by ID
// GetApiById 根据id获取api
func GetApiById(r *ghttp.Request) {
	var G request.GetById
	if err := r.Parse(&G); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	apiReturn, err := service.GetApiById(&G)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取失败, err:%v", err))
		r.Exit()
	}
	global.OkDetailed(r, g.Map{"api": apiReturn}, "获取成功")
}

// GetAllApis Gets all apis not paged
// GetAllApis 获取所有的Api 不分页
func GetAllApis(r *ghttp.Request) {
	apisReturn, err := service.GetAllApis()
	if err != nil {
		global.FailWithMessage(r, "获取失败")
		r.Exit()
	}
	global.OkDetailed(r, g.Map{"apis": apisReturn}, "获取成功")
}

// GetApiList Paging gets the LIST of apis
// GetApiList 分页获取API列表, 条件搜索后端看此api
func GetApiList(r *ghttp.Request) {
	var pageInfo request.GetApiList // 此结构体仅本方法使用
	if err := r.Parse(&pageInfo); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetApiInfoList(&pageInfo)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，%v", err))
		r.Exit()
	}
	global.OkDetailed(r, response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功")
}
