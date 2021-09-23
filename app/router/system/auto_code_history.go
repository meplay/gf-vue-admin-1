package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/api/system"
	"github.com/flipped-aurora/gf-vue-admin/interfaces"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var _ interfaces.Router = (*autoCodeHistory)(nil)

type autoCodeHistory struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewAutoCodeHistoryRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &autoCodeHistory{router: router, response: &response.Handler{}}
}

func (r *autoCodeHistory) Public() interfaces.Router {
	return r
}

func (r *autoCodeHistory) Private() interfaces.Router {
	return r
}

func (r *autoCodeHistory) PublicWithoutRecord() interfaces.Router {
	return r
}

func (r *autoCodeHistory) PrivateWithoutRecord() interfaces.Router {
	group := r.router.Group("/autoCode")
	{
		group.POST("getMeta", r.response.Handler()(system.AutoCodeHistory.First))         // 根据id获取meta信息
		group.POST("delSysHistory", r.response.Handler()(system.AutoCodeHistory.Delete))  // 删除回滚记录
		group.POST("rollback", r.response.Handler()(system.AutoCodeHistory.RollBack))     // 回滚
		group.POST("getSysHistory", r.response.Handler()(system.AutoCodeHistory.GetList)) // 获取回滚记录分页
	}
	return r
}
