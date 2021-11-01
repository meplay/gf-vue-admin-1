package email

import (
	"github.com/flipped-aurora/gf-vue-admin/app/router/middleware"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.PrivateRouter = (*_email)(nil)

type _email struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewEmailPrivateRouter(router *ghttp.RouterGroup) interfaces.PrivateRouter {
	return &_email{router: router, response: &response.Handler{}}
}

func (r *_email) Private() interfaces.PrivateRouter {
	group := r.router.Middleware(middleware.OperationRecord)
	{
		group.POST("emailTest", r.response.Handler()(Api.Test))
		group.POST("sendEmail", r.response.Handler()(Api.Send))
	}
	return r
}

func (r *_email) PrivateWithoutRecord() interfaces.PrivateRouter {
	return r
}
