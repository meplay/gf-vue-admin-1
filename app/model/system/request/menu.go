package request

import "github.com/flipped-aurora/gf-vue-admin/app/model/system"

type MenuCreate struct {
	system.Menu
}

type MenuUpdate struct {
	system.Menu
}

func (r *MenuUpdate) Update() map[string]interface{} {
	return map[string]interface{}{
		"path":         r.Path,
		"icon":         r.Icon,
		"sort":         r.Sort,
		"name":         r.Name,
		"title":        r.Title,
		"hidden":       r.Hidden,
		"parent_id":    r.ParentId,
		"close_tab":    r.CloseTab,
		"component":    r.Component,
		"keep_alive":   r.KeepAlive,
		"default_menu": r.DefaultMenu,
	}
}

type MenuAddAuthority struct {
	Menus       []system.Menu `json:"menus"`
	AuthorityId string        `json:"authorityId" example:"角色Id"`
}

func (r *MenuAddAuthority) ToAuthoritySetMenu() AuthoritySetMenu {
	entity := system.Authority{Menus: r.Menus, AuthorityId: r.AuthorityId}
	return AuthoritySetMenu{Authority: entity}
}
