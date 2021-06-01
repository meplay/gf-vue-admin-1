package service

import (
	"errors"
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/app/model/system/request"
	"gf-vue-admin/app/service/system/internal"
	"gf-vue-admin/library/global"
	"gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"gorm.io/gorm"
	"strconv"
)

var Authority = new(authority)

type authority struct {
	_menu             model.Menu
	_authority        model.Authority
	_dataAuthorities  model.DataAuthorities
	_authoritiesMenus model.AuthoritiesMenus
}

// Create 创建一个角色
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authority) Create(info *request.CreateAuthority) error {
	var authority *model.Authority
	if !errors.Is(global.Db.Where("authority_id = ?", info.AuthorityId).First(&model.Authority{}).Error, gorm.ErrRecordNotFound) {
		g.Log().Error(response.ErrorSameAuthorityId, g.Map{"authority_id": info.AuthorityId})
		return response.ErrorSameAuthorityId
	}
	authority = info.Create()
	err := global.Db.Create(&authority).Error
	return err
}

// First 根据角色id获取角色信息
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authority) First(info *request.GetAuthorityId) (result *model.Authority, err error) {
	var entity model.Authority
	err = global.Db.Where("authority_id = ?", info.AuthorityId).First(&entity).Error
	entity.DataAuthority = internal.Authority().GetDataAuthority(info.AuthorityId)
	return &entity, err
}

// Copy 复制一个角色
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authority) Copy(info *request.CopyAuthority) error {
	var entity model.Authority
	if !errors.Is(global.Db.Where("authority_id = ?", info.Authority.AuthorityId).First(&entity).Error, gorm.ErrRecordNotFound) {
		g.Log().Error(response.ErrorSameAuthorityId, g.Map{"authority_id": info.Authority.AuthorityId})
		return response.ErrorSameAuthorityId
	}
	info.Authority.Children = []model.Authority{}
	if menus, err := AuthorityMenu.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: info.OldAuthorityId}); err != nil {
		return err
	} else {
		_menus := make([]model.Menu, 0, len(*menus))
		for _, v := range *menus {
			id, _ := strconv.Atoi(v.MenuId)
			v.Menu.ID = uint(id)
			_menus = append(_menus, v.Menu)
		}
		info.Authority.Menus = _menus
	}
	if err := global.Db.Create(&info.Authority).Error; err != nil {
		return err
	}
	if len(info.Authority.Menus) > 0 {
		for _, m := range info.Authority.Menus {
			_entity := &model.AuthoritiesMenus{MenuId: m.ID, AuthorityId: info.Authority.AuthorityId}
			if err := global.Db.Create(&_entity).Error; err != nil {
				return err
			}
		}
	}
	var paths = Casbin.GetPolicyPath(info.OldAuthorityId)
	if err := Casbin.Update(&request.UpdateCasbin{AuthorityId: info.Authority.AuthorityId, CasbinInfos: paths}); err != nil {
		_ = a.Delete(&request.GetAuthorityId{AuthorityId: info.Authority.AuthorityId})
	}

	for _, data := range info.Authority.DataAuthority {
		insert := &model.DataAuthorities{
			AuthorityId:   info.Authority.AuthorityId,
			DataAuthority: data.AuthorityName,
		}
		if err := global.Db.Create(&insert).Error; err != nil {
			return err
		}
	}
	return nil
}

// Update 更改一个角色
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authority) Update(info *request.UpdateAuthority) error {
	return global.Db.Where("authority_id", info.AuthorityId).Updates(info.Update()).Error
}

// Delete 删除角色
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authority) Delete(info *request.GetAuthorityId) error {
	var entity model.Authority
	if !errors.Is(global.Db.Where("parent_id = ?", info.AuthorityId).First(&entity).Error, gorm.ErrRecordNotFound) {
		return response.ErrorHasSonAuthority
	}
	var user model.Admin
	if !errors.Is(global.Db.Where("authority_id = ?",info.AuthorityId).First(&user).Error, gorm.ErrRecordNotFound) {
		return response.ErrorUseAuthority
	}
	if err := global.Db.Where("authority_id = ?",info.AuthorityId).First(&entity).Error; err != nil {
		return err
	}
	entity.Menus = internal.Authority().GetMenus(entity.AuthorityId)
	entity.DataAuthority = internal.Authority().GetDataAuthority(entity.AuthorityId)
	if err := global.Db.Unscoped().Delete(&entity, info.AuthorityId).Error; err != nil {
		return err
	}
	if len(entity.Menus) > 0 {
		if err := global.Db.Delete(&model.AuthoritiesMenus{}, entity.AuthorityId).Error; err != nil {
			return err
		}
	}
	if len(entity.DataAuthority) > 0 {
		if err := global.Db.Delete(&model.DataAuthorities{}, entity.AuthorityId).Error; err != nil {
			return err
		}
	}
	Casbin.ClearCasbin(0, info.AuthorityId)
	return nil
}

// GetList 分页获取数据
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authority) GetList(info *request.PageInfo) (list interface{}, total int64, err error) {
	var authorities []model.Authority
	db := global.Db.Model(&model.Authority{})
	err = db.Count(&total).Error
	err = db.Scopes(internal.Gorm.Paginate(info)).Find(&authorities).Error
	if len(authorities) > 0 {
		for i, b := range authorities {
			authorities[i].DataAuthority = internal.Authority().GetDataAuthority(b.AuthorityId)
			internal.Authority().FindChildren(&authorities[i])
		}
	}
	return authorities, total, err
}

// SetDataAuthority 设置角色资源权限
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authority) SetDataAuthority(info *request.SetDataAuthority) error {
	if err := global.Db.Unscoped().Delete(&model.DataAuthorities{}, info.AuthorityId).Error; err != nil {
		return err
	}
	for _, d := range info.DataAuthorityId {
		entity := &model.DataAuthorities{AuthorityId: info.AuthorityId, DataAuthority: d.AuthorityId}
		if err := global.Db.Create(&entity).Error; err != nil {
			return err
		}
	}
	return nil
}

// SetMenuAuthority 菜单与角色绑定
// Author [Aizen1172](https://github.com/Aizen1172)
func (a *authority) SetMenuAuthority(info *model.Authority) error {
	var entity model.Authority
	if err := global.Db.Where("authority_id", info.AuthorityId).First(&entity).Error; err != nil {
		return err
	}
	entity.Menus = info.Menus
	return internal.Authority().ReplaceMenu(&entity)
}
