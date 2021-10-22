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
		group.POST("create",c.response.Handler()(example.Customer.Create))		//创建客户表
		group.POST("update",c.response.Handler()(example.Customer.Update))		//更新{.Description}
		group.POST("delete",c.response.Handler()(example.Customer.Delete))		// 删除{.Description}}
		group.DELETE("deletes",c.response.Handler()(example.Customer.Deletes))	// 批量删除{.Description}}
	}
	return c
}

func (c *customer) PublicWithoutRecord() interfaces.Router {
	return c
}

func (c *customer) PrivateWithoutRecord() interfaces.Router {
	group := c.router.Group("/customer")
	{
		group.GET("first",c.response.Handler()(example.Customer.First))			// 根据id获取{.Description}}
		group.POST("getList", c.response.Handler()(example.Customer.GetList))	// 分页获取{.Description}}列表
	}
	return c
}

