package internal

import (
	model "gf-vue-admin/app/model/system"
	"github.com/gogf/gf/frame/g"
)

var Authority = new(authority)

func init() {
	entities := make([]model.Authority, 0, 1)
	if err := g.DB().Table(entities[0].TableName()).Structs(&entities); err != nil {
		g.Log().Error("获取全部 Authority 失败!", g.Map{"err": err})
	} else {
		_map := make(map[string][]model.Authority, len(entities))
		for _, entity := range entities {
			Authority.authorityMap[entity.AuthorityId] = entity
			if entity.ParentId != "0" || &entity != nil {
				if value, ok := _map[entity.ParentId]; ok {
					value = append(value, entity)
					_map[entity.ParentId] = value
				} else {
					var a1 = make([]model.Authority, 0, 1)
					a1 = append(a1, entity)
					_map[entity.ParentId] = a1
				}
			}
		}
		Authority.authoritiesMap = _map
	}
}

type authority struct {
	authorityMap   map[string]model.Authority
	authoritiesMap map[string][]model.Authority
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 查询资源角色
func (a *authority) GetDataAuthority(id string) (result *[]model.Authority) {
	var entities = make([]model.DataAuthority, 0, 10)
	if err := g.DB().Table(entities[0].TableName()).Where(g.Map{"authority_id": id}).Struct(&entities); err != nil {
		g.Log().Error("查询角色的资源角色失败!", g.Map{"err": err})
		return nil
	}
	var authorities = make([]model.Authority, 0, len(entities))
	for _, entity := range entities {
		var a1 = a.authorityMap[entity.DataAuthority]
		authorities = append(authorities, a1)
	}
	return &authorities
}


//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 查询子角色
func (a *authority) FindChildren(authority *model.Authority) {
	authority.Children = a.authoritiesMap[authority.AuthorityId]
	if len(authority.Children) > 0 {
		for i := range authority.Children {
			a.FindChildren(&authority.Children[i])
		}
	}
}
