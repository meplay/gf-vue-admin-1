package internal

import (
	"flipped-aurora/gf-vue-admin/server/app/model/system/request"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Context = new(_context)

type _context struct{}

// GetAdminID 从GoFrame的Context中获取从jwt解析出来的用户ID
// Author: [SliverHorn](https://github.com/SliverHorn)
func (i *_context) GetAdminID(r *ghttp.Request) uint {
	var claims request.CustomClaims
	param := r.GetParam("claims")
	if err := gconv.Struct(param, &claims); err != nil {
		g.Log().Error("Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件! ", g.Map{"err": err})
		r.ExitAll()
	}
	return claims.AdminId
}

// GetUserUuid 从GoFrame的Context中获取从jwt解析出来的用户UUID
// Author: [SliverHorn](https://github.com/SliverHorn)
func (i *_context) GetUserUuid(r *ghttp.Request) string {
	var claims request.CustomClaims
	param := r.GetParam("claims")
	if err := gconv.Struct(param, &claims); err != nil {
		g.Log().Error("Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件! ", g.Map{"err": err})
		r.ExitAll()
	}
	return claims.AdminUuid
}

// GetUserAuthorityId 从GoFrame的Context中获取从jwt解析出来的AuthorityId
// Author: [SliverHorn](https://github.com/SliverHorn)
func (i *_context) GetUserAuthorityId(r *ghttp.Request) string {
	var claims request.CustomClaims
	param := r.GetParam("claims")
	if err := gconv.Struct(param, &claims); err != nil {
		g.Log().Error("Context中获取从jwt解析出来的AuthorityId失败, 请检查路由是否使用jwt中间件! ", g.Map{"err": err})

		g.Log().Error("获取jwt中间件信息失败!", g.Map{"err": err})
		r.ExitAll()
	}
	return claims.AdminAuthorityId
}

// GetAdminClaims 获取jwt里含有的管理员信息
// Author: [SliverHorn](https://github.com/SliverHorn)
func (i *_context) GetAdminClaims(r *ghttp.Request) *request.CustomClaims {
	var claims request.CustomClaims
	param := r.GetParam("claims")
	if err := gconv.Struct(param, &claims); err != nil {
		g.Log().Error("Context中获取jwt里含有的管理员信息, 请检查路由是否使用jwt中间件! ", g.Map{"err": err})
		r.ExitAll()
	}
	return &claims
}
