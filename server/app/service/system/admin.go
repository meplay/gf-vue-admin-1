package service

import (
	"errors"
	model "flipped-aurora/gf-vue-admin/server/app/model/system"
	"flipped-aurora/gf-vue-admin/server/app/model/system/request"
	"flipped-aurora/gf-vue-admin/server/app/service/system/internal"
	"flipped-aurora/gf-vue-admin/server/library/global"
	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

// Register 注册
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *admin) Register(info *request.Register) error {
	var entity model.Admin
	err := global.Db.Where("username = ?", info.Username).First(&entity).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册")
	}
	entity = info.Create()
	if err = entity.EncryptedPassword(); err != nil {
		return err
	}
	return global.Db.Create(&entity).Error
}

// Login 登录
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *admin) Login(info *request.AdminLogin) (result *model.Admin, err error) {
	var entity model.Admin
	if errors.Is(global.Db.Where("username = ?", info.Username).Preload("Authority").First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	}
	if !entity.CompareHashAndPassword(info.Password) {
		return &entity, errors.New("密码错误")
	}
	return &entity, nil
}

// First 根据ID获取用户信息
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *admin) First(info *request.GetById) (result *model.Admin, err error) {
	var entity model.Admin
	err = global.Db.First(&entity, info.Id).Preload("Authority").Error
	return &entity, err
}

// FirstByUuid 根据uuid获取用户信息
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *admin) FirstByUuid(uuid string) (result *model.Admin, err error) {
	var entity model.Admin
	err = global.Db.First(&entity, "uuid = ?", uuid).Error
	return &entity, err
}

// SetUserInfo 更新用户信息
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *admin) SetUserInfo(info *request.UpdateAdmin) (result *model.Admin, err error) {
	entity := info.Update()
	err = global.Db.Model(&model.Admin{}).Where("uuid = ?", info.Uuid).Updates(&entity).Error
	return &entity, err
}

// ChangePassword 修改密码
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *admin) ChangePassword(info *request.ChangePassword) error {
	var entity model.Admin
	err := global.Db.Where("username = ?", info.Username).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("用户不存在")
	}
	if !entity.CompareHashAndPassword(info.Password) {
		return errors.New("密码错误")
	}
	entity.Password = info.NewPassword
	if err = entity.EncryptedPassword(); err != nil {
		return err
	}
	return global.Db.Where("username = ?", info.Username).Update("password", entity.Password).Error
}

// SetUserAuthority 设置用户角色
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *admin) SetUserAuthority(info *request.SetAuthority) error {
	return global.Db.Model(&model.Admin{}).Where("uuid = ?", info.Uuid).Update("authority_id", info.AuthorityId).Error
}

// Delete 删除用户
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *admin) Delete(info *request.GetById) error {
	return global.Db.Delete(&model.Admin{}, info.Id).Error
}

// GetList 获取用户列表
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *admin) GetList(info *request.PageInfo) (list *[]model.Admin, total int64, err error) {
	var entities []model.Admin
	db := global.Db.Model(&model.Admin{})
	err = db.Count(&total).Error
	err = db.Scopes(internal.Gorm.Paginate(info)).Preload("Authority").Find(&entities).Error
	return &entities, total, err
}
