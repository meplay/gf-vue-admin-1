package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/app/router/middleware"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*authority)(nil)

type authority struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewAuthorityRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &authority{router: router, response: &response.Handler{}}
}

func (r *authority) Public() interfaces.Router {
	return r
}

func (r *authority) Private() interfaces.Router {
	group := r.router.Group("/authority").Middleware(middleware.OperationRecord)
	{
		group.POST("createAuthority", r.response.Handler()(system.Authority.Create))                 // 创建角色
		group.POST("copyAuthority", r.response.Handler()(system.Authority.Copy))                     // 更新角色
		group.PUT("updateAuthority", r.response.Handler()(system.Authority.Update))                  // 更新角色
		group.POST("deleteAuthority", r.response.Handler()(system.Authority.Delete))                 // 删除角色
		group.POST("setDataAuthority", r.response.Handler()(system.Authority.SetAuthorityResources)) // 设置角色资源权限
	}
	return r
}

func (r *authority) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *authority) PrivateWithoutRecord() interfaces.Router {
	group := r.router.Group("/authority")
	{
		group.POST("getAuthorityList", r.response.Handler()(system.Authority.GetList)) // 获取角色列表
	}
	return r
}
