package email

import (
	"github.com/flipped-aurora/gf-vue-admin/app/router/middleware"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*_email)(nil)

type _email struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewEmailRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &_email{router: router, response: &response.Handler{}}
}

func (r *_email) Public() interfaces.Router {
	return r
}

func (r *_email) Private() interfaces.Router {
	group := r.router.Middleware(middleware.OperationRecord)
	{
		group.POST("emailTest", r.response.Handler()(Api.Test))
		group.POST("sendEmail", r.response.Handler()(Api.Send))
	}
	return r
}

func (r *_email) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *_email) PrivateWithoutRecord() interfaces.Router {
	return r
}
