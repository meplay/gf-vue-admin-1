package v1

import (
	"fmt"
	"server/app/api/request"
	"server/app/api/response"
	"server/app/service"
	"server/library/global"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

// UpdateCasbin Change role API permissions
// UpdateCasbin 更改角色api权限
func UpdateCasbin(r *ghttp.Request) {
	var update request.CasbinInReceive
	if err := r.Parse(&update); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.UpdateCasbin(update.AuthorityId, update.CasbinInfos); err != nil {
		global.OkWithMessage(r, fmt.Sprintf("添加规则失败，%v", err))
		r.Exit()
	}
	global.OkWithMessage(r, "添加规则成功")
}

// GetPolicyPathByAuthorityId Get permission list
// GetPolicyPathByAuthorityId 获取权限列表
func GetPolicyPathByAuthorityId(r *ghttp.Request) {
	var cmr request.CasbinInReceive
	if err := r.Parse(&cmr); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	paths := service.GetPolicyPathByAuthorityId(cmr.AuthorityId)
	global.OkWithData(r, response.PolicyPath{Paths: paths})
}

// CasbinTest casb RBAC RESTFUL测试路由
func CasbinTest(r *ghttp.Request) {
	// 测试restful以及占位符代码  随意书写
	pathParam := r.GetParam("pathParam")
	query := r.GetQuery("query")
	global.OkDetailed(r, g.Map{"pathParam": pathParam, "query": query}, "获取规则成功")
}
