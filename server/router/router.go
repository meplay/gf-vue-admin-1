package router

import "github.com/gogf/gf/frame/g"

// InitializeRouters 初始化总路由
func InitializeRouters() {
	InitBaseRouter()             // 初始化基础功能路由 不做鉴权
	InitAdminsRouter()           // 初始化用户路由
	InitMenuRouter()             // 初始化menu路由
	InitAuthorityRouter()        // 初始化角色路由
	InitApiRouter()              // 初始化功能api路由
	InitCasbinRouter()           // 初始化权限相关路由
	InitJwtRouter()              // 初始化jwt相关路由
	InitAutoCodeRouter()         // 创建自动化代码
	InitOperationRouter()        // 初始化操作记录路由
	InitDictionaryRouter()       // 初始化字典管理路由
	InitDictionaryDetailRouter() // 初始化字典详情管理路由
	InitCustomerRouter()         // 初始化客户路由
	InitFileRouter()             // 初始化文件上传下载功能路由
	// InitWorkflowRouter()         // 初始化工作流相关路由
	// InitSystemRouter()           // 初始化system配置相关路由
	g.Log().Info(g.I18n().Translate(`{#Prefix} {#router} {#register} {#success}`, "zh-CN"))
}
