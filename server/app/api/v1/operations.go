package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/net/ghttp"
)

// CreateOperation Create Operation
// CreateOperation 创建Operation
func CreateOperation(r *ghttp.Request) {
	var c request.CreateOperation
	if err := r.Parse(&c); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.CreateOperation(&c); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("创建失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "创建成功")
}

// DeleteOperation Delete Operation
// DeleteOperation 删除Operation
func DeleteOperation(r *ghttp.Request) {
	var d request.DeleteOperation
	if err := r.Parse(&d); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteOperation(&d); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// DeleteOperations Batch delete Operation
// DeleteOperations 批量删除Operation
func DeleteOperations(r *ghttp.Request) {
	var d request.DeleteOperations
	if err := r.Parse(&d); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteOperations(&d); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("批量删除失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "批量删除成功")
}

// UpdateOperation Update Operation
// UpdateOperation 更新Operation
func UpdateOperation(r *ghttp.Request) {
	var u request.UpdateOperation
	if err := r.Parse(&u); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.UpdateOperation(&u); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("更新失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "更新成功")
}

// FindOperation Query Operation with id
// FindOperation 用id查询Operation
func FindOperation(r *ghttp.Request) {
	var u request.UpdateOperation
	if err := r.Parse(&u); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.UpdateOperation(&u); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "获取成功")
}

// GetOperationList Page out the Operation list
// GetOperationList 分页获取Operation列表
func GetOperationList(r *ghttp.Request) {
	var g request.GetOperationList
	if err := r.Parse(&g); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetOperationList(&g)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，%v", err))
		r.Exit()
	}
	global.OkWithData(r, response.PageResult{
		List:     list,
		Total:    total,
		Page:     g.Page,
		PageSize: g.PageSize,
	})
}
