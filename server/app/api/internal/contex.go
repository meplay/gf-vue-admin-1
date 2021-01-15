package internal

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"server/app/api/request"
)

var Info = new(info)

type info struct {
	claims *request.CustomClaims
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 从GoFrame的Context中获取从jwt解析出来的用户ID
func (i *info) GetAdminID(r *ghttp.Request) uint {
	var claims = r.GetParam("claims")
	if err := gconv.Struct(claims, i.claims); err != nil {
		g.Log().Errorf("管理员信息失败!, err:%v", err)
		r.ExitAll()
	}
	return i.claims.AdminId
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 从GoFrame的Context中获取从jwt解析出来的用户UUID
func (i *info) GetUserUuid(r *ghttp.Request) string {
	var claims = r.GetParam("claims")
	if err := gconv.Struct(claims, i.claims); err != nil {
		g.Log().Errorf("管理员信息失败!, err:%v", err)
		r.ExitAll()
	}
	return i.claims.AdminUuid
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 从GoFrame的Context中获取从jwt解析出来的用户UUID
func (i *info) GetUserAuthorityId(r *ghttp.Request) string {
	var claims = r.GetParam("claims")
	if err := gconv.Struct(claims, i.claims); err != nil {
		g.Log().Errorf("管理员信息失败!, err:%v", err)
		r.ExitAll()
	}
	return i.claims.AdminAuthorityId
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取jwt里含有的管理员信息
func (i *info) GetAdminClaims(r *ghttp.Request)  *request.CustomClaims {
	var claims = r.GetParam("claims")
	if err := gconv.Struct(claims, claims); err != nil {
		g.Log().Errorf("管理员信息失败!, err:%v", err)
		r.ExitAll()
	}
	return i.claims
}
