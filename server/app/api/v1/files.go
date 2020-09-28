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

// UploadFile Upload file example
// UploadFile 上传文件示例
func UploadFile(r *ghttp.Request) {
	noSave := r.GetQueryString("noSave", "0")
	_, header, err := r.Request.FormFile("file")
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("上传文件失败，%v", err))
		r.Exit()
	}
	file, uploadErr := service.UploadFile(header, noSave)
	if uploadErr != nil {
		global.FailWithMessage(r, fmt.Sprintf("修改数据库链接失败，%v", err))
		r.Exit()
	}
	global.OkDetailed(r, g.Map{"file": file}, "上传成功")
}

// DeleteFile Delete File
// DeleteFile 删除文件
func DeleteFile(r *ghttp.Request) {
	var d request.DeleteFile
	if err := r.Parse(&d); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.DeleteFile(&d); err != nil {
		global.FailWithMessage(r, fmt.Sprintf("删除失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "删除成功")
}

// GetFileList List of paging files
// GetFileList 分页文件列表
func GetFileList(r *ghttp.Request) {
	var g request.PageInfo
	if err := r.Parse(&g); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	list, total, err := service.GetFileList(&g)
	if err != nil {
		global.FailWithMessage(r, fmt.Sprintf("获取数据失败，%v", err))
		r.Exit()
	}
	global.OkDetailed(r, response.PageResult{List: list, Total: total, Page: g.Page, PageSize: g.PageSize}, "获取成功")
}
