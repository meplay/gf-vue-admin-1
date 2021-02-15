package service

import (
	"database/sql"
	"errors"
	"gf-vue-admin/app/api/request"
	"gf-vue-admin/app/api/response"
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/app/service/system/internal"
	"github.com/gogf/gf/frame/g"
)

var Authority = new(authority)

type authority struct {
	_authority model.Authority
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 创建一个角色
func (a *authority) Create(info *request.CreateAuthority) error {
	if !errors.Is(g.DB().Table(a._authority.TableName()).Where("authority_id = ?", info.AuthorityId).Struct(&model.Authority{}), sql.ErrNoRows) {
		g.Log().Error(response.ErrorSameAuthorityId, g.Map{"authority_id": info.AuthorityId})
		return response.ErrorSameAuthorityId
	}
	_, err := g.DB().Table(a._authority.TableName()).Insert(info.Create())
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 复制一个角色
func (a *authority) Copy(info *request.CopyAuthority) error {
	var entity model.Authority
	if !errors.Is(g.DB().Table(a._authority.TableName()).Where("authority_id = ?", info.Authority.AuthorityId).Struct(&entity), sql.ErrNoRows) {
		g.Log().Error(response.ErrorSameAuthorityId, g.Map{"authority_id": info.Authority.AuthorityId})
		return response.ErrorSameAuthorityId
	}
	//info.Authority.Children = []model.Authority{}
	//if a.menus, err = a.authority.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: info.OldAuthorityId}); err != nil {
	//	return nil, errors.New("GetMenuAuthority Failed! err: " + err.Error())
	//}
	//for _, v := range *a.menus {
	//	id, _ := strconv.Atoi(v.MenuId)
	//	v.Menu.ID = uint(id)
	//	a.baseMenu = append(a.baseMenu, v.Menu)
	//}
	//info.Authority.BaseMenus = a.baseMenu
	//err = global.Db.Create(&info.Authority).Error
	//
	//var paths = a.casbin.GetPolicyPathByAuthorityId(info.OldAuthorityId)
	//if err = a.casbin.Update(info.Authority.AuthorityId, paths); err != nil {
	//	_ = a.Delete(&request.DeleteAuthority{GetAuthorityId: request.GetAuthorityId{AuthorityId: info.Authority.AuthorityId}})
	//}
	return nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 更改一个角色
func (a *authority) Update(info *request.UpdateAuthority) error {
	_, err := g.DB().Table(a._authority.TableName()).Update(info.Update(), g.Map{"authority_id": info.AuthorityId})
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 删除角色
func (a *authority) Delete(info *request.GetAuthorityId) error {
	var entity model.Authority
	if !errors.Is(g.DB().Table(a._authority.TableName()).Where(info.Condition()).Struct(&entity), sql.ErrNoRows) {
		return response.ErrorUseAuthority
	}
	if !errors.Is(g.DB().Table(a._authority.TableName()).Where("parent_id = ?", info.AuthorityId).Struct(&entity), sql.ErrNoRows) {
		return response.ErrorHasSonAuthority
	}
	//var db = global.Db.Preload("BaseMenus").Preload("DataAuthority").Where("authority_id = ?", info.AuthorityId).First(&entity)
	//var err = db.Unscoped().Delete(&entity).Error
	//if len(entity.BaseMenus) > 0 {
	//	err = global.Db.Model(&entity).Association("BaseMenus").Delete(&entity.BaseMenus)
	//}
	//if len(entity.DataAuthority) > 0 {
	//	err = global.Db.Model(&entity).Association("DataAuthority").Delete(&entity.DataAuthority)
	//}
	//a.casbin.ClearCasbin(0, info.AuthorityId)
	return nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 分页获取数据
func (a *authority) GetList(info *request.PageInfo) (list interface{}, total int, err error) {
	var authorities []model.Authority
	db := g.DB().Table(a._authority.TableName())
	limit, offset := info.Paginate()
	total, err = db.Where(g.Map{"parent_id": "0"}).Count()
	err = db.Limit(limit).Offset(offset).Where(g.Map{"parent_id": "0"}).Structs(&authorities)
	if len(authorities) > 0 {
		for i, b := range authorities {
			authorities[i].DataAuthority = *internal.Authority.GetDataAuthority(b.AuthorityId)
			internal.Authority.FindChildren(&authorities[i])
		}
	}
	return authorities, total, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 设置角色资源权限
func (a *authority) SetDataAuthority(info *request.SetDataAuthority) error {
	if _, err := g.DB().Table(a._authority.TableName()).Delete(g.Map{"authority_id": info.AuthorityId}); err != nil {
		return err
	}
	for _, d := range info.DataAuthorityId {
		entity := &model.DataAuthority{AuthorityId: info.AuthorityId, DataAuthority: d.AuthorityId}
		if _, err := g.DB().Table(entity.TableName()).Insert(entity); err != nil {
			return err
		}
	}
	return nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 菜单与角色绑定
func (a *authority) SetMenuAuthority(info *model.Authority) error {
	var entity model.DataAuthority
	err := g.DB().Table(a._authority.TableName()).Where(g.Map{"authority_id": info.AuthorityId}).Struct(&entity)
	// todo menu 与 authority 数据联动
	//var err = global.Db.Preload("BaseMenus").First(&entity, " = ?", info.AuthorityId).Error
	////err = global.Db.Model(&entity).Association("menus").Replace(&entity.menus)
	//err = global.Db.Model(&model.Authority{}).Association("BaseMenus").Replace(&info.BaseMenus)
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 获取所有角色信息
func (a *authority) First(info *request.GetAuthorityId) (result *model.Authority, err error) {
	var entity model.Authority
	entity.DataAuthority = *internal.Authority.GetDataAuthority(info.AuthorityId)
	return &entity, err
}
