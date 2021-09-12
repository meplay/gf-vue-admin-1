package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system/internal"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

// GetMenu
// @Tags SystemAuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body common.Empty true "空"
// @Success 200 {object} response.Response{data=[]system.AuthorityMenu} "获取列表数据成功!"
// @Router /menu/getMenu [post]
func (a *authorityMenu) GetMenu(r *ghttp.Request) *response.Response {
	claims := internal.NewClaims(r)
	data, err := system.AuthorityMenu.GetMenuTree(claims.GetUserUuid())
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: g.Map{"menus": data}, MessageCode: response.SuccessGetList}
}


// GetMenuAuthority
// @Tags SystemAuthorityMenu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetAuthorityId true "角色ID"
// @Success 200 {object} response.Response{data=[]system.AuthorityMenu} "获取列表数据成功!"
// @Router /menu/getMenuAuthority [post]
func (a *authorityMenu) GetMenuAuthority(r *ghttp.Request) *response.Response {
	var info common.GetAuthorityId
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	data, err := system.AuthorityMenu.GetAuthorityMenu(&info)
	if err != nil {
		return &response.Response{Data: g.Map{"menus":data}, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: g.Map{"menus":data}, MessageCode: response.SuccessGetList}
}