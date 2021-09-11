package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
)

var Authority = new(authority)

type authority struct{}

// Create 创建一个角色
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) Create(info *request.AuthorityCreate) (data *system.Authority, err error) {
	var entity system.Authority
	if !errors.Is(global.Db.Where("authority_id = ?", info.AuthorityId).First(&entity).Error, gorm.ErrRecordNotFound) {
		return &entity, errors.New("存在相同角色id")
	}
	entity = info.Create()
	err = global.Db.Create(&entity).Error
	return &entity, err
}

// Copy 复制一个角色
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) Copy(info *request.AuthorityCopy) (*system.Authority, error) {
	var entity system.Authority
	if !errors.Is(global.Db.Where("authority_id = ?", info.Authority.AuthorityId).First(&entity).Error, gorm.ErrRecordNotFound) {
		return &entity, errors.New("存在相同角色id")
	}
	info.Authority.Children = []system.Authority{}
	menus, err := AuthorityMenu.GetAuthorityMenu(&common.GetAuthorityId{AuthorityId: info.OldAuthorityId})
	var menu []system.Menu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.Menu.ID = uint(intNum)
		menu = append(menu, v.Menu)
	}
	info.Authority.Menus = menu
	if err = global.Db.Create(&info.Authority).Error; err != nil {
		return nil, errors.Wrap(err, "复制角色失败!")
	}

	paths := Casbin.GetPolicyPathByAuthorityId(info.OldAuthorityId)
	if err = Casbin.Update(info.Authority.AuthorityId, paths); err != nil {
		if err = s.Delete(&request.AuthorityDelete{Authority: info.Authority}); err != nil {
			return nil, errors.Wrap(err, "更新权限失败, 删除角色失败!")
		}
	}
	return &info.Authority, nil
}

// Update 更新角色
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) Update(info *request.AuthorityUpdate) (authority *system.Authority, err error) {
	err = global.Db.Model(&system.Authority{}).Where("authority_id = ?", info.AuthorityId).Updates(&info.Authority).Error
	return &info.Authority, err
}

// Delete 删除角色
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) Delete(info *request.AuthorityDelete) error {
	if !errors.Is(global.Db.Where("authority_id = ?", info.AuthorityId).First(&system.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.Db.Where("parent_id = ?", info.AuthorityId).First(&system.Authority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := global.Db.Preload("Menus").Where("authority_id = ?", info.AuthorityId).First(info)
	err := db.Unscoped().Delete(info).Error
	if len(info.Menus) > 0 {
		err = global.Db.Model(info).Association("Menus").Delete(info.Menus)
	} else {
		err = db.Error
	}
	err = global.Db.Delete(&[]system.UseAuthority{}, "authority_id = ?", info.AuthorityId).Error
	Casbin.Clear(0, info.AuthorityId)
	return err
}

// GetAuthorityInfo 获取角色所有信息
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) GetAuthorityInfo(info common.GetAuthorityId) (*system.Authority, error) {
	var entity system.Authority
	err := global.Db.Preload("AuthorityResources").Where("authority_id = ?", info.AuthorityId).First(&entity).Error
	return &entity, err
}

// SetAuthorityResources 设置角色资源权限
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) SetAuthorityResources(info *request.AuthoritySetResources) error {
	var entity system.Authority
	if err := global.Db.Preload("AuthorityResources").First(&entity, "authority_id = ?", info.AuthorityId).Error; err != nil {
		return errors.Wrap(err, "角色查找失败!")
	}
	if err := global.Db.Model(&entity).Association("AuthorityResources").Replace(&info.AuthorityResources); err != nil {
		return errors.Wrap(err, "设置角色资源权限!")
	}
	return nil
}

// SetAuthorityMenu 菜单与角色绑定
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) SetAuthorityMenu(auth *request.AuthoritySetMenu) error {
	var entity system.Authority
	if err := global.Db.Preload("Menus").First(&entity, "authority_id = ?", auth.AuthorityId).Error; err != nil {
		return errors.Wrap(err, "菜单查找失败!")
	}
	if err := global.Db.Model(&entity).Association("Menus").Replace(&auth.Menus); err != nil {
		return errors.Wrap(err, "设置菜单与角色绑定!")
	}
	return nil
}

// GetList 分页获取数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) GetList(info *common.PageInfo) (list []system.Authority, total int64, err error) {
	var entities []system.Authority
	db := global.Db.Model(&system.Authority{})
	err = db.Scopes(common.Paginate(info)).Preload("AuthorityResources").Where("parent_id = 0").Find(&entities).Error
	length := len(entities)
	for i := 0; i < length; i++ {
		if err = s.findChildrenAuthority(&entities[i]); err != nil {
			return entities, total, err
		}
	}
	return entities, total, err
}

// findChildrenAuthority 查询子角色 todo 循环sql优化
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) findChildrenAuthority(authority *system.Authority) (err error) {
	err = global.Db.Preload("AuthorityResources").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = s.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
