package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system/internal"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/auth"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"strconv"
)

var User = new(user)

type user struct{}

// Register
// @Tags SystemUser
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body request.UserRegister true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {object} response.Response{} "注册成功!"
// @Router /user/register [post]
func (a *user) Register(r *ghttp.Request) *response.Response {
	var info request.UserRegister
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	data, err := system.User.Register(&info)
	if err != nil {
		return &response.Response{Data: g.Map{"user": data}, Error: err, Message: "注册失败!"}
	}
	return &response.Response{Data: g.Map{"user": data}, Message: "注册成功!"}
}

// Login
// @Tags SystemUser
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.UserLogin true "请求参数"
// @Success 200 {object} response.Response{data=response.UserLogin} "登录成功!"
// @Router /base/login [post]
func (a *user) Login(r *ghttp.Request) *response.Response {
	var info request.UserLogin
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if system.Store.Verify(info.CaptchaId, info.Captcha, true) {
		data, err := system.User.Login(&info)
		if err != nil {
			return &response.Response{Error: err, Message: "登录失败!"}
		}
		return &response.Response{Data: data, Message: "登录成功!"}
	} else {
		return &response.Response{Code: 7, Message: "验证码错误!"}
	}
}

// GetUserInfo
// @Tags SystemUser
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Success 200 {object} response.Response{data=system.User} "获取用户信息成功!"
// @Router /user/getUserInfo [get]
func (a *user) GetUserInfo(r *ghttp.Request) *response.Response {
	var info request.UserFind
	claims := internal.NewClaims(r)
	if info.Uuid = claims.GetUserUuid(); info.Uuid == "" || claims.Error() != nil {
		return &response.Response{Error: claims.Error(), Message: "获取用户信息失败!"}
	}
	data, err := system.User.Find(&info)
	if err != nil {
		return &response.Response{Error: err, Message: "获取用户信息失败!"}
	}
	return &response.Response{Data: g.Map{"userInfo": data}, Message: "获取用户信息成功!"}
}

// SetUserInfo
// @Tags SystemUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UserUpdate true "请求参数"
// @Success 200 {object} response.Response{data=system.User} "设置成功!"
// @Router /user/setUserInfo [put]
func (a *user) SetUserInfo(r *ghttp.Request) *response.Response {
	var info request.UserUpdate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	data, err := system.User.Update(&info)
	if err != nil {
		return &response.Response{Error: err, Message: "设置失败!"}
	}
	return &response.Response{Data: g.Map{"userInfo": data}, Message: "设置成功!"}
}

// ChangePassword
// @Tags SystemUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.UserChangePassword true "请求参数"
// @Success 200 {object} response.Response{} "修改成功!"
// @Router /user/changePassword [post]
func (a *user) ChangePassword(r *ghttp.Request) *response.Response {
	var info request.UserChangePassword
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := system.User.ChangePassword(&info); err != nil {
		return &response.Response{Error: err, Message: "修改失败!"}
	}
	return &response.Response{Message: "修改成功!"}
}

// Delete
// @Tags SystemUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "删除成功!"
// @Router /user/deleteUser [delete]
func (a *user) Delete(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	claims := internal.NewClaims(r)
	if id := claims.GetUserID(); id == 0 || claims.Error() != nil {
		if id == info.ID {
			return &response.Response{Error: claims.Error(), Message: "自我删除失败!"}
		}
	}
	if err := system.User.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// GetList
// @Tags SystemUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.PageInfo true "请求参数"
// @Success 200 {object} response.Response{data=[]system.User} "获取列表数据成功!"
// @Router /user/getUserList [post]
func (a *user) GetList(r *ghttp.Request) *response.Response {
	var info common.PageInfo
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := system.User.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: common.NewPageResult(list, total, info), MessageCode: response.SuccessGetList}
}

// SetUserAuthority
// @Tags SystemUser
// @Summary 更改用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UserSetAuthority true "请求参数"
// @Success 200 {object} response.Response{} "修改成功!"
// @Router /user/setUserAuthority [post]
func (a *user) SetUserAuthority(r *ghttp.Request) *response.Response {
	var info request.UserSetAuthority
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	claims := internal.NewClaims(r)
	if info.ID = claims.GetUserID(); info.ID == 0 || claims.Error() != nil {
		err := claims.Error()
		return &response.Response{Error: err, Message: err.Error()}
	}
	if info.Uuid = claims.GetUserUuid(); info.Uuid == "" || claims.Error() != nil {
		err := claims.Error()
		return &response.Response{Error: err, Message: err.Error()}
	}
	if err := system.User.SetAuthority(&info); err != nil {
		return &response.Response{Error: err, Message: "修改失败!"}
	}
	_claims := claims.GetUserClaims()
	_claims.AuthorityId = info.AuthorityId
	if token, err := auth.NewJWT().CreateToken(_claims); err != nil {
		return &response.Response{Error: err, Message: "修改失败!"}
	} else {
		r.Response.Header().Set("new-token", token)
		r.Response.Header().Set("new-expires-at", strconv.FormatInt(_claims.ExpiresAt, 10))
		return &response.Response{Message: "修改成功!"}
	}
}

// SetUserAuthorities
// @Tags SystemUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UserSetAuthorities true "请求参数"
// @Success 200 {object} response.Response{} "修改成功!"
// @Router /user/setUserAuthorities [post]
func (a *user) SetUserAuthorities(r *ghttp.Request) *response.Response {
	var info request.UserSetAuthorities
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := system.User.SetUserAuthorities(&info); err != nil {
		return &response.Response{Error: err, Message: "修改失败!"}
	}
	return &response.Response{Message: "修改成功!"}
}
