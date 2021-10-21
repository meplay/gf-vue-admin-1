package example

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/example"
	"github.com/flipped-aurora/gf-vue-admin/app/router/middleware"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*file)(nil)

type file struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewFileRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &file{router: router, response: &response.Handler{}}
}

func (r *file) Public() interfaces.Router {
	return r
}

func (r *file) Private() interfaces.Router {
	group := r.router.Group("/fileUploadAndDownload").Middleware(middleware.OperationRecord)
	{
		group.POST("upload", r.response.Handler()(example.File.Upload))       // 上传文件
		group.POST("deleteFile", r.response.Handler()(example.File.Delete))   // 删除指定文件
		group.POST("getFileList", r.response.Handler()(example.File.GetList)) // 获取上传文件列表
	}
	return r
}

func (r *file) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *file) PrivateWithoutRecord() interfaces.Router {
	return r
}
