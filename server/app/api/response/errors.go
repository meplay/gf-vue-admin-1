package response

import (
	"errors"
)

var (
	// Admin
	ErrorUserNoExist        = errors.New(`用户不存在! err: record not found`)
	ErrorAdminDelete        = errors.New(`用户名为admin, 不允许删除`)
	ErrorWrongPassword      = errors.New(`密码错误`)
	ErrorEncryptedPassword  = errors.New(`密码加密失败`)
	ErrorUsernameRegistered = errors.New(`用户名已注册`)

	// Jwt
	ErrorJwt2Redis     = errors.New(`设置登录状态失败`)
	ErrorCreateToken   = errors.New(`创建token失败`)
	ErrorInvalidateJwt = errors.New(`jwt作废失败`)

	// Authority
	ErrorUseAuthority    = errors.New(`此角色有用户正在使用禁止删除`)
	ErrorHasSonAuthority = errors.New(`此角色存在子角色不允许删除`)
	ErrorSameAuthorityId = errors.New("已存在相同的角色id")

	// Dictionary
	ErrorSameType = errors.New("已存在数据,不允许创建")

	// Api
	ErrorSameApi = errors.New("已存在记录")

	// JsonWebToken
	TokenExpired     = errors.New(`无法处理这个令牌! `)
	TokenInvalid     = errors.New(`令牌已过期! `)
	TokenMalformed   = errors.New(`这不是一个有效的令牌! `)
	TokenNotValidYet = errors.New(`令牌尚未激活! `)

	// Menu
	ErrorUpdateMenu       = errors.New(`更新失败`)
	ErrorUpdateMenuName   = errors.New(`存在相同name, 修改失败`)
	ErrorHasChildrenMenu  = errors.New(`此菜单存在子菜单不可删除`)
	ErrorCreateParameters = errors.New(`创建menu的Parameters失败`)

	// Casbin
	ErrorAddPolicies = errors.New(`存在相同api, 添加失败, 请联系管理员`)
)
