package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"github.com/gogf/gf/net/ghttp"
)

type AuthorityRouter struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewAuthorityRouter(router *ghttp.RouterGroup) *AuthorityRouter {
	return &AuthorityRouter{router: router, response: &response.Handler{}}
}

func (a *AuthorityRouter) Init() {
	var authority = a.router.Group("/authority")
	{
		authority.POST("createAuthority", a.response.Handler()(api.Authority.Create))            // 创建角色
		authority.POST("copyAuthority", a.response.Handler()(api.Authority.Copy))                // 拷贝角色
		authority.PUT("updateAuthority", a.response.Handler()(api.Authority.Update))             // 更新角色
		authority.POST("deleteAuthority", a.response.Handler()(api.Authority.Delete))            // 删除角色
		authority.POST("getAuthorityList", a.response.Handler()(api.Authority.GetList))          // 获取角色列表
		authority.POST("setDataAuthority", a.response.Handler()(api.Authority.SetDataAuthority)) // 设置角色资源权限
	}
}
