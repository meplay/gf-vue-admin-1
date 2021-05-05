package service

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/app/model/system/request"
	"gf-vue-admin/library/global"
)

type AdminInterface interface {
	Register(info *request.Register) error
	Login(info *request.AdminLogin) (result *model.Admin, err error)
	First(info *request.GetById) (result *model.Admin, err error)
	FindAdmin(info *request.GetByUuid) (result *model.Admin, err error)
	Update(info *request.UpdateAdmin) (result *model.Admin, err error)
	ChangePassword(info *request.ChangePassword) error
	SetUserAuthority(info *request.SetAuthority) error
	Delete(info *request.GetById) error
	GetList(info *request.PageInfo) (list interface{}, total int, err error)
}

func Admin() AdminInterface {
	switch global.Config.System.OrmType {
	case "gdb":
		return AdminGdb
	case "gorm":
		return AdminGorm
	default:
		return AdminGdb
	}
}
