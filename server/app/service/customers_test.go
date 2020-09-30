package service

import (
	"fmt"
	"server/app/model/customers"
	"testing"

	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/frame/g"
)

func TestFindCustomers(t *testing.T) {
	var err error
	var data *customers.CustomerHasOneAdmin
	data = (*customers.CustomerHasOneAdmin)(nil)
	db := g.DB("default").Table("customers").Safe()
	adminDb := g.DB("default").Table("admins").Safe()
	err = db.Where(g.Map{"id": 1}).Struct(&data)
	err = adminDb.Where(g.Map{"id": data.Id}).Struct(&data.Admin)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

func TestGetCustomersList(t *testing.T) {
	var err error
	datalist := ([]*customers.CustomerHasOneAdmin)(nil)
	db := g.DB("default").Table("customers").Safe()
	adminDb := g.DB("default").Table("admins").Safe()
	err = db.ScanList(&datalist, "Customers")
	err = adminDb.
		Where("id", gdb.ListItemValues(datalist, "Customers", "SysUserId")).
		ScanList(&datalist, "Admin", "Customers", "id:Id")
	if err != nil {
		panic(err)
	}
	fmt.Println(datalist)
}
