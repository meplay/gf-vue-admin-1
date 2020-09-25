package service

import (
	"fmt"
	"server/app/model/admins"
	"server/library/global"
	"testing"

	"github.com/gogf/gf/frame/g"
)

func TestAdminLogin(t *testing.T) {
	var admin admins.AdminHasOneAuthority
	err := g.DB(global.Db).Table("admins").Scan(&admin, g.Map{"username": "admin"})
	if err != nil {
		panic(err)
	}
	err = g.DB(global.Db).Table("admins").Scan(&admin.Authority, g.Map{"authority_id": admin.AuthorityId})
	if err != nil {
		panic(err)
	}
	fmt.Println(admin)
}
