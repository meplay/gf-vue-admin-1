package router

import (
	"gf-vue-admin/app/api/response"
	api "gf-vue-admin/app/api/system"
	"gf-vue-admin/interfaces"
	"github.com/gogf/gf/net/ghttp"
)

type JwtBlacklist struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewJwtBlacklistGroup(router *ghttp.RouterGroup) interfaces.Router {
	return &JwtBlacklist{router: router, response: &response.Handler{}}
}

func (j *JwtBlacklist) Init() {
	var jwt = j.router.Group("/jwt")
	{
		jwt.POST("jsonInBlacklist", j.response.Handler()(api.JwtBlacklist.JwtToBlacklist)) // jwt加入黑名单
	}
}
