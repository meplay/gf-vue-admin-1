package router

import "github.com/gogf/gf/frame/g"

// InitializeRouters 初始化总路由
func InitializeRouters() {
	InitJwtRouter()              // jwt相关路由
	InitApiRouter()              // 功能api路由
	InitFileRouter()             // 文件上传下载功能路由
	InitMenuRouter()             // menu路由
	InitBaseRouter()             // 基础功能路由 不做鉴权
	InitEmailRouter()            // 邮件相关路由
	InitAdminsRouter()           // 用户路由
	InitCasbinRouter()           // 权限相关路由
	InitSystemRouter()           // 系统配置路由
	InitCustomerRouter()         // 客户路由
	InitAutoCodeRouter()         // 创建自动化代码
	InitAuthorityRouter()        // 角色路由
	InitOperationRouter()        // 操作记录路由
	InitDictionaryRouter()       // 字典管理路由
	InitSimpleUploadRouter()     // 断点续传功能路由
	InitDictionaryDetailRouter() // 字典详情管理路由
	g.Log().Info(g.I18n().Translate(`{#Prefix} {#router} {#register} {#success}`, "zh-CN"))
}
