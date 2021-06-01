package request

import (
	model "gf-vue-admin/app/model/system"
)

type UpdateMenu struct {
	model.Menu
}

func (u *UpdateMenu) Update() *model.Menu {
	return &model.Menu{
		Path:      u.Path,
		Name:      u.Name,
		ParentId:  u.ParentId,
		Component: u.Component,
		Sort:      u.Sort,
		Hidden:    u.Hidden,
		Meta: model.Meta{
			KeepAlive:   u.KeepAlive,
			DefaultMenu: u.DefaultMenu,
			Title:       u.Title,
			Icon:        u.Icon,
		},
	}
}
