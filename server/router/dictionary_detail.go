package router

import (
	v1 "server/app/api/v1"
	"server/app/middleware"

	"github.com/gogf/gf/frame/g"
)

// InitDictionaryDetailRouter 注册字典详情管理路由
func InitDictionaryDetailRouter() {
	DictionaryDetailRouter := g.Server().Group("sysDictionaryDetail").Middleware(
		middleware.JwtAuth,
		middleware.CasbinMiddleware,
	)
	{
		DictionaryDetailRouter.POST("createSysDictionaryDetail", v1.CreateDictionaryDetail)   // 新建DictionaryDetail
		DictionaryDetailRouter.DELETE("deleteSysDictionaryDetail", v1.DeleteDictionaryDetail) // 删除DictionaryDetail
		DictionaryDetailRouter.PUT("updateSysDictionaryDetail", v1.UpdateDictionaryDetail)    // 更新DictionaryDetail
		DictionaryDetailRouter.GET("findSysDictionaryDetail", v1.FindDictionaryDetail)        // 根据ID获取DictionaryDetail
		DictionaryDetailRouter.GET("getSysDictionaryDetailList", v1.GetDictionaryDetailList)  // 获取DictionaryDetail列表
	}
}
