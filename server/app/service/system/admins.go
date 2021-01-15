package service

import (
	"gf-vue-admin/app/api/request"
	"gf-vue-admin/app/api/response"
	"gf-vue-admin/app/model/system"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

var Admin = &admin{db: g.DB().Table("admins").Safe()}

type admin struct {
	db *gdb.Model
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 修改管理员密码
func (a *admin) ChangePassword(info *request.ChangePassword) error {
	var entity = (*model.Admin)(nil)
	if err := a.db.Where("username", info.Username).Struct(&entity); err != nil {
		return response.ErrorUserNoExist
	}
	if !entity.CompareHashAndPassword(info.Password) {
		return response.ErrorWrongPassword
	}
	entity.Password = info.NewPassword
	if err := entity.EncryptedPassword(); err != nil {
		return err
	}
	if _, err := a.db.Where("username", info.Username).Data(g.Map{"password": entity.Password}).Update(); err != nil {
		return err
	}
	return nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取用户列表
func (a *admin) GetAdminList(info *request.PageInfo) (list interface{}, total int, err error) {
	var admins = ([]*model.Admin)(nil)
	var db = a.db.Safe()
	limit, offset := info.Paginate()
	total, err = db.Count()
	err = db.Limit(limit).Offset(offset).Structs(&admins)
	return admins, total, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 用于刷新token,根据uuid返回admin信息
func (a *admin) FindAdmin(info *request.GetByUuid) (result *model.Admin, err error) {
	var entity = (*model.Admin)(nil)
	var db = a.db.Safe()
	err = db.Where(g.Map{"uuid": info.Uuid}).Struct(&entity)
	return entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 用于刷新token,根据uuid返回admin信息
func (a *admin) FindAdminById(info *request.GetById) (admin *model.Admin, err error) {
	var entity = (*model.Admin)(nil)
	err = a.db.Where(g.Map{"id": info.Id}).Struct(&entity)
	return entity, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 设置用户权限
func (a *admin) SetUserAuthority(set *request.SetAuthority) error {
	var _, err = a.db.Update(g.Map{"authority_id": set.AuthorityId}, g.Map{"uuid": set.Uuid})
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除管理员
func (a *admin) Delete(delete *request.GetById) (err error) {
	_, err = a.db.Delete(g.Map{"id": delete.Id})
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 设置管理员信息
func (a *admin) SetAdminInfo(info *request.UpdateAdmin) (admin *model.Admin, err error) {
	_, err = a.db.Update(g.Map{"avatar": info.Avatar}, g.Map{"uuid": info.Uuid})
	return a.FindAdmin(&request.GetByUuid{Uuid: info.Uuid})
}
