package data

import (
	"gf-vue-admin/interfaces"
	extra "gf-vue-admin/library/data/extra"
	system "gf-vue-admin/library/data/system"
	workflow "gf-vue-admin/library/data/workflow"
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
