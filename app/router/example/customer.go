package example

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/example"
	"github.com/flipped-aurora/gf-vue-admin/app/router/middleware"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*customer)(nil)

type customer struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewCustomerRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &customer{router: router, response: &response.Handler{}}
}

func (c *customer) Public() interfaces.Router {
	return c
}

func (c *customer) Private() interfaces.Router {
	group := c.router.Group("/customer").Middleware(middleware.OperationRecord)
	{
		group.POST("customer", c.response.Handler()(example.Customer.Create))   // 创建客户
		group.PUT("customer", c.response.Handler()(example.Customer.Update))    // 更新客户
		group.DELETE("customer", c.response.Handler()(example.Customer.Delete)) // 删除客户
	}
	return c
}

func (c *customer) PublicWithoutRecord() interfaces.Router {
	return c
}

func (c *customer) PrivateWithoutRecord() interfaces.Router {
	group := c.router.Group("/customer")
	{
		group.GET("customer", c.response.Handler()(example.Customer.First))       // 获取单一客户信息
		group.GET("customerList", c.response.Handler()(example.Customer.GetList)) // 获取客户列表
	}
	return c
}
