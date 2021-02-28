package router

import (
	api "gf-vue-admin/app/api/extra"
	"gf-vue-admin/app/api/response"
	"gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type file struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewFileRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &file{router: router, response: &response.Handler{}}
}

func (f *file) Init() {
	group := f.router.Group("/fileUploadAndDownload")
	{
		group.POST("/upload", f.response.Handler()(api.File.UploadFile))   // 上传文件
		group.POST("/deleteFile", f.response.Handler()(api.File.Delete))   // 删除指定文件
		group.POST("/getFileList", f.response.Handler()(api.File.GetList)) // 获取上传文件列表
	}
}
