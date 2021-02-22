package router

import (
	api "gf-vue-admin/app/api/extra"
	"gf-vue-admin/app/api/response"
	"gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type FileRouter struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewFileRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &FileRouter{router: router, response: &response.Handler{}}
}

func (f *FileRouter) Init() {
	var file = f.router.Group("/fileUploadAndDownload")
	{
		file.POST("/upload", f.response.Handler()(api.File.UploadFile))   // 上传文件
		file.POST("/deleteFile", f.response.Handler()(api.File.Delete))   // 删除指定文件
		file.POST("/getFileList", f.response.Handler()(api.File.GetList)) // 获取上传文件列表
	}
}
