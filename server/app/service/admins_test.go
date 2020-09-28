package service

import (
	"fmt"
	"server/app/model/admins"
	"testing"

	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/frame/g"
)

func TestGetAdminList(t *testing.T) {
	var adminList []admins.AdminHasOneAuthority
	err := g.DB("default").Table("admins").ScanList(&adminList, "AdminTest")
	if err != nil {
		panic(err)
	}
	err = g.DB("default").Table("admins").
		Where("authority_id", gdb.ListItemValues(adminList, "AdminTest", "AuthorityId")).
		ScanList(&adminList, "Authority", "AdminTest", "authority_id:AuthorityId")
	if err != nil {
		panic(err)
	}
	fmt.Println(adminList)
}
