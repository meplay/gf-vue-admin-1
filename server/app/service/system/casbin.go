package service

import (
	model "flipped-aurora/gf-vue-admin/server/app/model/system"
	"flipped-aurora/gf-vue-admin/server/app/model/system/request"
	"flipped-aurora/gf-vue-admin/server/library/gdbadapter"
	"flipped-aurora/gf-vue-admin/server/library/global"
	"flipped-aurora/gf-vue-admin/server/library/response"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	"github.com/gogf/gf/frame/g"
	"strings"
)

var Casbin = new(_casbin)

type _casbin struct{}

// Update 更新casbin权限
// Author: [SliverHorn](https://github.com/SliverHorn)
func (c *_casbin) Update(info *request.UpdateCasbin) error {
	c.ClearCasbin(0, info.AuthorityId)
	rules := make([][]string, 0, len(info.CasbinInfos))
	for _, v := range info.CasbinInfos {
		entity := model.Casbin{PType: "p", AuthorityId: info.AuthorityId, Path: v.Path, Method: v.Method}
		rules = append(rules, []string{entity.AuthorityId, entity.Path, entity.Method})
	}
	e := c.Casbin()
	if success, _ := e.AddPolicies(rules); success == false {
		return response.ErrorAddPolicies
	}
	return nil
}

// UpdateApi API更新随动
// Author [Aizen1172](https://github.com/Aizen1172)
func (c *_casbin) UpdateApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	entity := &model.Casbin{Path: newPath, Method: newMethod}
	err := global.Db.Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(entity).Error
	return err
}

// GetPolicyPath 获取权限列表
// Author: [SliverHorn](https://github.com/SliverHorn)
func (c *_casbin) GetPolicyPath(authorityId string) (pathMaps []request.CasbinInfo) {
	enforcer := c.Casbin()
	list := enforcer.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{Path: v[1], Method: v[2]})
	}
	return pathMaps
}

// ClearCasbin 清除匹配的权限
// Author: [SliverHorn](https://github.com/SliverHorn)
func (c *_casbin) ClearCasbin(v int, p ...string) bool {
	e := c.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

// Clear API清除
// Author [Aizen1172](https://github.com/Aizen1172)
func (c *_casbin) Clear(path, method string) bool {
	var rule model.Casbin
	if err := global.Db.Delete(&rule, path, method).Error; err != nil {
		return false
	}
	return true
}

// Casbin 持久化到数据库  引入自定义规则
// Author: [SliverHorn](https://github.com/SliverHorn)
func (c *_casbin) Casbin() *casbin.Enforcer {
	a, _ := gdbadapter.NewAdapterByDB(g.DB(), "casbin_rule")
	e, _ := casbin.NewEnforcer(global.Config.Casbin.ModelPath, a)
	e.AddFunction("ParamsMatch", c.ParamsMatchFunc)
	_ = e.LoadPolicy()
	return e
}

// ParamsMatch 自定义规则函数
// Author: [SliverHorn](https://github.com/SliverHorn)
func (c *_casbin) ParamsMatch(key1 string, key2 string) bool {
	key1 = strings.Split(key1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// ParamsMatchFunc 自定义规则函数
// Author: [SliverHorn](https://github.com/SliverHorn)
func (c *_casbin) ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	var name1 = args[0].(string)
	var name2 = args[1].(string)
	return c.ParamsMatch(name1, name2), nil
}
