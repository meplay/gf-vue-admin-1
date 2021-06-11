package router

import (
	api "flipped-aurora/gf-vue-admin/server/app/api/extra"
	"flipped-aurora/gf-vue-admin/server/interfaces"
	"flipped-aurora/gf-vue-admin/server/library/response"
	"github.com/gogf/gf/net/ghttp"
)

type uploader struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewSimpleUploaderRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &uploader{router: router, response: &response.Handler{}}
}

func (s *uploader) Init() {
	group := s.router.Group("/simpleUploader")
	{
		group.POST("upload", s.response.Handler()(api.SimpleUploader.Upload))            // 上传功能
		group.GET("checkFileMd5", s.response.Handler()(api.SimpleUploader.CheckFileMd5)) // 文件完整度验证
		group.GET("mergeFileMd5", s.response.Handler()(api.SimpleUploader.MergeFileMd5)) // 合并文件
	}
}
