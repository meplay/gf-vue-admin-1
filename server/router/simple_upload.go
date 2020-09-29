package router

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitSimpleUploadRouter 注册断点续传路由
func InitSimpleUploadRouter() {
	SimpleUploadRouter := g.Server().Group("simpleUploader").Middleware(
		middleware.JwtAuth,
	)
	{
		SimpleUploadRouter.POST("upload", v1.CreateSimpleUpload) // 上传功能
		SimpleUploadRouter.GET("checkFileMd5", v1.CheckFileMd5)  // 文件完整度验证
		SimpleUploadRouter.GET("mergeFileMd5", v1.MergeFileMd5)  // 合并文件
	}
}
