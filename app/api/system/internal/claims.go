package internal

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/gogf/gf/net/ghttp"
	_errors "github.com/pkg/errors"
	"sync"
)

var (
	claimsInstance *_claims
	//mutex 互斥锁
	mutex sync.Mutex
)

type _claims struct {
	err    error
	data   *request.CustomClaims
	isInit bool
}

func NewClaims(r *ghttp.Request) *_claims {
	mutex.Lock()
	defer mutex.Unlock()
	if claimsInstance == nil {
		claimsInstance = &_claims{}
		claimsInstance.Initialize(r)
		claimsInstance.isInit = true
	}
	return claimsInstance
}

func (i *_claims) Initialize(r *ghttp.Request) {
	data := r.GetCtxVar("claims")
	var claims request.CustomClaims
	if err := data.Struct(&claims); err != nil {
		i.err = _errors.New("反射失败!")
	}
	i.data = &claims
}

// GetUserID 从GoFrame的Context中获取从jwt解析出来的用户ID
// Author: [SliverHorn](https://github.com/SliverHorn)
func (i *_claims) GetUserID() uint {
	if !i.isInit {
		i.err = _errors.New("获取用户的ID失败, 请检查路由是否使用jwt中间件!")
		return 0
	}
	return i.data.ID
}

// GetUserUuid 从GoFrame的Context中获取从jwt解析出来的用户UUID
// Author: [SliverHorn](https://github.com/SliverHorn)
func (i *_claims) GetUserUuid() string {
	if !i.isInit {
		i.err = _errors.New("获取用户的uuid失败, 请检查路由是否使用jwt中间件!")
		return ""
	}
	return i.data.Uuid
}

// GetUserAuthorityId 从GoFrame的Context中获取从jwt解析出来的AuthorityId
// Author: [SliverHorn](https://github.com/SliverHorn)
func (i *_claims) GetUserAuthorityId() string {
	if !i.isInit {
		i.err = _errors.New("获取用户的角色id失败, 请检查路由是否使用jwt中间件!")
		return ""
	}
	return i.data.AuthorityId
}

// GetUserClaims 获取jwt里含有的用户信息
// Author: [SliverHorn](https://github.com/SliverHorn)
func (i *_claims) GetUserClaims() *request.CustomClaims {
	if !i.isInit {
		i.err = _errors.New("获取jwt里含有的用户信息失败, 请检查路由是否使用jwt中间件!")
		return nil
	}
	return i.data
}

func (i *_claims) Error() error {
	return i.err
}
