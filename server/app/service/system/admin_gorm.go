package service

import (
	"errors"
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/app/model/system/request"
	"gf-vue-admin/app/service/system/internal"
	"gf-vue-admin/library/global"
	"gf-vue-admin/library/response"
	"gorm.io/gorm"
)

var AdminGorm = new(adminGorm)

type adminGorm struct{}

// Register 注册逻辑代码
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *adminGorm) Register(info *request.Register) error {
	if errors.Is(global.Db.Where("username = ?", info.Username).First(&model.Admin{}).Error, gorm.ErrRecordNotFound) {
		return response.ErrorUsernameRegistered
	}
	entity := info.Create()
	if err := entity.EncryptedPassword(); err != nil {
		return response.ErrorEncryptedPassword
	}
	return global.Db.Create(&entity).Error
}

// Login @author: [SliverHorn](https://github.com/SliverHorn)
//@description: 设置管理员信息
func (a *adminGorm) Login(info *request.AdminLogin) (result *model.Admin, err error) {
	var entity model.Admin
	if errors.Is(global.Db.Where("username = ?", info.Username).Preload("Authority").First(&entity).Error, gorm.ErrRecordNotFound) {
		return &entity, response.ErrorUserNoExist
	}
	if !entity.CompareHashAndPassword(info.Password) {
		return &entity, response.ErrorWrongPassword
	}
	return &entity, nil
}

// First 用于刷新token,根据uuid返回admin信息
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *adminGorm) First(info *request.GetById) (result *model.Admin, err error) {
	var entity model.Admin
	err = global.Db.First(&entity, info.Id).Error
	return &entity, err
}

// FindAdmin 用于刷新token,根据uuid返回admin信息
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *adminGorm) FindAdmin(info *request.GetByUuid) (result *model.Admin, err error) {
	var entity model.Admin
	err = global.Db.Find(&entity, "uuid = ?", info.Uuid).Error
	return &entity, err
}

// Update 设置管理员信息
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *adminGorm) Update(info *request.UpdateAdmin) (result *model.Admin, err error) {
	var entity model.Admin
	if errors.Is(global.Db.Where("uuid = ?", info.Uuid).First(&entity).Error, gorm.ErrRecordNotFound) {
		return &entity, response.ErrorUserNoExist
	}
	entity.Avatar = info.Avatar
	err = global.Db.Select("avatar").Where("uuid = ?", info.Uuid).Updates(&entity).Error
	return &entity, err
}

// ChangePassword 修改管理员密码
// Author: [SliverHorn](https://github.com/SliverHorn)
func (a *adminGorm) ChangePassword(info *request.ChangePassword) error {
	var entity model.Admin
	if errors.Is(global.Db.Where("username = ?", info.Username).Preload("Authority").First(&entity).Error, gorm.ErrRecordNotFound) {
		return response.ErrorUserNoExist
	}
	if !entity.CompareHashAndPassword(info.Password) {
		return response.ErrorWrongPassword
	}
	entity.Password = info.NewPassword
	if err := entity.EncryptedPassword(); err != nil {
		return err
	}
	return global.Db.Where("username = ?", info.Username).Update("password", entity.Password).Error
}

// SetUserAuthority 设置用户权限
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *adminGorm) SetUserAuthority(info *request.SetAuthority) error {
	return global.Db.Where("uuid = ?", info.Uuid).Update("authority_id", info.AuthorityId).Error
}

// Delete 删除管理员
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *adminGorm) Delete(info *request.GetById) error {
	return global.Db.Delete(&model.Admin{}, info.Id).Error
}

// GetList 分页获取用户列表
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *adminGorm) GetList(info *request.PageInfo) (list interface{}, total int, err error) {
	var (
		t        int64
		entities []model.Admin
	)
	db := global.Db.Model(&model.Admin{})
	err = db.Count(&t).Error
	err = db.Scopes(internal.Gorm.Paginate(info)).Preload("Authority").Find(&entities).Error
	return entities, int(t), err
}
