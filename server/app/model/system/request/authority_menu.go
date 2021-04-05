package request

import model "gf-vue-admin/app/model/system"

type AddMenuAuthority struct {
	GetAuthorityId
	Menus       []model.Menu
}
