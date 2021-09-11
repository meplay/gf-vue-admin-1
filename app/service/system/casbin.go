package system

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"strings"
	"sync"
)

var (
	once           sync.Once
	Casbin         = new(_casbin)
	syncedEnforcer *casbin.SyncedEnforcer
)

type _casbin struct{}

// Casbin 持久化到数据库  引入自定义规则
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_casbin) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := adapter.NewAdapterByDB(global.Db)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.Config.Casbin.ModelPath, a)
		syncedEnforcer.AddFunction("ParamsMatch", s.ParamsMatchFunc)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

func (s *_casbin) Update(authorityId string, casbinInfos []request.CasbinInfo) error {
	s.Clear(0, authorityId)
	length := len(casbinInfos)
	rules := make([][]string, 0, length)
	for _, v := range casbinInfos {
		cm := system.Casbin{
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.AuthorityId, cm.Path, cm.Method})
	}
	e := s.Casbin()
	success, _ := e.AddPolicies(rules)
	if success == false {
		return errors.New("存在相同api,添加失败,请联系管理员!")
	}
	return nil
}

// UpdateApi API更新随动
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_casbin) UpdateApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.Db.Table("casbin_rule").Model(&system.Casbin{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// GetPolicyPathByAuthorityId 获取权限列表
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_casbin) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := s.Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// ClearCasbin 清除匹配的权限
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_casbin) Clear(v int, p ...string) bool {
	e := s.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

// ParamsMatch 自定义规则函数
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_casbin) ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0] // 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// ParamsMatchFunc 自定义规则函数
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_casbin) ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return s.ParamsMatch(name1, name2), nil
}
