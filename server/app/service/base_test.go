package service

import (
	"fmt"
	"server/app/model/admins"
	"testing"

	"github.com/gogf/gf/frame/g"
)

func TestAdminLogin(t *testing.T) {
	var admin admins.AdminHasOneAuthority
	var err error
	err = g.DB("default").Table("admins").Scan(&admin, g.Map{"username": "admin"})
	err = g.DB("default").Table("authorities").Scan(&admin.Authority, g.Map{"authority_id": admin.AuthorityId})
	if err != nil {
		panic(err)
	}
	fmt.Println(admin)
}
