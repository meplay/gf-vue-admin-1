package service

import (
	"database/sql"
	"errors"
	"gf-vue-admin/app/model/system"
	"gf-vue-admin/app/model/system/request"
	"gf-vue-admin/app/service/system/internal"
	"gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
)

var AdminGdb = new(adminGdb)

type adminGdb struct {
	_admin model.Admin
}

// Register 注册逻辑代码
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *adminGdb) Register(info *request.Register) error {
	var entity model.Admin
	if !errors.Is(g.DB().Table(a._admin.TableName()).Where(g.Map{"username": info.Username}).Struct(&entity), sql.ErrNoRows) {
		return response.ErrorUsernameRegistered
	}
	user := info.Create()
	if err := user.EncryptedPassword(); err != nil {
		return response.ErrorEncryptedPassword
	}
	_, err := g.DB().Table(a._admin.TableName()).Data(user).Insert()
	return err
}

// Login @author: [SliverHorn](https://github.com/SliverHorn)
//@description: 设置管理员信息
func (a *adminGdb) Login(info *request.AdminLogin) (result *model.Admin, err error) {
	var entity model.Admin
	if err = g.DB().Table(a._admin.TableName()).Where(g.Map{"username": info.Username}).Struct(&entity); err != nil {
		return &entity, response.ErrorUserNoExist
	}
	if !entity.CompareHashAndPassword(info.Password) {
		return &entity, response.ErrorWrongPassword
	}
	entity.Authority = internal.Authority().First(entity.AuthorityId)
	return &entity, nil
}

// First 用于刷新token,根据uuid返回admin信息
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *adminGdb) First(info *request.GetById) (result *model.Admin, err error) {
	var entity model.Admin
	err = g.DB().Table(a._admin.TableName()).Where(g.Map{"id": info.Id}).Struct(&entity)
	return &entity, err
}

// Update 设置管理员信息
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *adminGdb) Update(info *request.UpdateAdmin) (result *model.Admin, err error) {
	_, err = g.DB().Table(a._admin.TableName()).Update(g.Map{"avatar": info.Avatar}, g.Map{"uuid": info.Uuid})
	return a.FindAdmin(&request.GetByUuid{Uuid: info.Uuid})
}

// ChangePassword 修改管理员密码
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *adminGdb) ChangePassword(info *request.ChangePassword) error {
	var entity model.Admin
	if err := g.DB().Table(a._admin.TableName()).Where("username", info.Username).Struct(&entity); err != nil {
		return response.ErrorUserNoExist
	}
	if !entity.CompareHashAndPassword(info.Password) {
		return response.ErrorWrongPassword
	}
	entity.Password = info.NewPassword
	if err := entity.EncryptedPassword(); err != nil {
		return err
	}
	if _, err := g.DB().Table(a._admin.TableName()).Where("username", info.Username).Data(g.Map{"password": entity.Password}).Update(); err != nil {
		return err
	}
	return nil
}

// SetUserAuthority 设置用户权限
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *adminGdb) SetUserAuthority(info *request.SetAuthority) error {
	_, err := g.DB().Table(a._admin.TableName()).Update(g.Map{"authority_id": info.AuthorityId}, g.Map{"uuid": info.Uuid})
	return err
}

// FindAdmin 用于刷新token,根据uuid返回admin信息
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *adminGdb) FindAdmin(info *request.GetByUuid) (result *model.Admin, err error) {
	var entity model.Admin
	db := g.DB().Table(a._admin.TableName()).Safe()
	err = db.Where(g.Map{"uuid": info.Uuid}).Struct(&entity)
	return &entity, err
}

// Delete 删除管理员
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *adminGdb) Delete(info *request.GetById) error {
	_, err := g.DB().Table(a._admin.TableName()).Delete(info.Condition())
	return err
}

// GetList 分页获取用户列表
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *adminGdb) GetList(info *request.PageInfo) (list interface{}, total int, err error) {
	var admins []model.Admin
	db := g.DB().Table(a._admin.TableName()).Safe()
	limit, offset := info.Paginate()
	total, err = db.Count()
	err = db.Limit(limit).Offset(offset).Structs(&admins)
	for i, entity := range admins {
		admins[i].Authority = internal.Authority().First(entity.AuthorityId)
	}
	return admins, total, err
}
