package data

import (
	"flipped-aurora/gf-vue-admin/server/interfaces"
	extra "flipped-aurora/gf-vue-admin/server/library/data/extra"
	system "flipped-aurora/gf-vue-admin/server/library/data/system"
	workflow "flipped-aurora/gf-vue-admin/server/library/data/workflow"
)

func Initialize() error {
	system.Init()
	return interfaces.InitDb(
		system.Api,
		system.Menu,
		system.Admin,
		system.Casbin,
		system.Authority,
		system.Dictionary,
		system.AuthorityMenu,
		system.AuthoritiesMenus,
		system.ResourcesAuthorities,
		system.DictionaryDetail,

		workflow.Edge,
		workflow.Node,
		workflow.EndPoint,
		workflow.StartPoint,
		workflow.Process,

		extra.File,
	)
}
