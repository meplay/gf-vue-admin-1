package request

import model "flipped-aurora/gf-vue-admin/server/app/model/system"

type AddMenuAuthority struct {
	GetAuthorityId
	Menus       []model.Menu
}
