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
	err = global.Db.Transaction(func(tx *gorm.DB) error {
		entity = info.Create()
		err = tx.Create(&entity).Error
		menus := info.DefaultMenu()
		if err = s.SetAuthorityMenuByTransaction(tx, menus); err != nil {
			return errors.Wrap(err, "新角色添加默认菜单权限失败!")
		}
		entities := info.DefaultCasbin()
		for i := 0; i < len(entities); i++ {
			entities[i].AuthorityId = info.AuthorityId
		}
		if err = tx.Create(&entities).Error; err != nil {
			return errors.Wrap(err, "新角色添加默认api权限失败!")
		}
		entity.Menus = menus.Menus
		return nil
	})
	return &entity, err
}

// Copy 复制一个角色
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) Copy(info *request.AuthorityCopy) (*system.Authority, error) {
	var entity system.Authority
	if !errors.Is(global.Db.Where("authority_id = ?", info.Authority.AuthorityId).First(&entity).Error, gorm.ErrRecordNotFound) {
		return &entity, errors.New("存在相同角色id")
	}

	menus, err := AuthorityMenu.GetAuthorityMenu(info.ToGetAuthorityId())
	if err != nil {
		return nil, errors.Wrap(err, "查询角色id("+info.OldAuthorityId+")的菜单数据失败")
	}

	length := len(menus)
	info.Authority.Menus = make([]system.Menu, 0, length)
	for i := 0; i < length; i++ {
		intNum, _ := strconv.Atoi(menus[i].MenuId)
		menus[i].Menu.ID = uint(intNum)
		info.Authority.Menus = append(info.Authority.Menus, menus[i].Menu)
	} // 角色菜单数据组装

	if err = global.Db.Create(&info.Authority).Error; err != nil {
		return nil, errors.Wrap(err, "复制角色失败!")
	}

	var casbinRules []system.Casbin
	casbinRules, _, err = Casbin.GetList(info.ToCasbinSearch())
	if err != nil {
		return nil, errors.Wrap(err, "查询角色id("+info.OldAuthorityId+")的权限数据失败")
	}
	length = len(casbinRules)
	for i := 0; i < length; i++ {
		casbinRules[i].AuthorityId = info.Authority.AuthorityId
	} // 替换旧角色id为复制后的角色id

	err = global.Db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&info.Authority).Error; err != nil {
			return errors.Wrap(err, "复制角色及角色菜单数据失败!")
		}
		if err = tx.Create(&casbinRules).Error; err != nil {
			return errors.Wrap(err, "复制角色权限数据失败!")
		}
		return nil
	})
	return &info.Authority, nil
}

// Update 更新角色
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) Update(info *request.AuthorityUpdate) (authority *system.Authority, err error) {
	updates := info.Update()
	err = global.Db.Select("*").Omit("created_at", "deleted_at").Updates(&updates).Error
	return &updates, err
}

// Delete 删除角色
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) Delete(info *request.AuthorityDelete) error {
	if !errors.Is(global.Db.Where("authority_id = ?", info.AuthorityId).First(&system.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除!")
	}
	if !errors.Is(global.Db.Where("parent_id = ?", info.AuthorityId).First(&system.Authority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除!")
	}

	if err := global.Db.Model(&system.Authority{}).Preload("Menus").Where("authority_id = ?", info.AuthorityId).First(&info.Authority).Error; err != nil {
		return errors.Wrap(err, "角色不存在!")
	} // 查询角色关联的菜单

	list, _, err := Casbin.GetList(info.ToCasbinSearch())
	if err != nil {
		return errors.Wrap(err, "查询角色关联的权限列表失败!")
	} // 查询角色关联的权限列表

	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&system.Authority{}).Unscoped().Delete(&info.Authority).Error; err != nil {
			return errors.Wrap(err, "删除角色失败!")
		} // 删除角色信息
		if len(info.Menus) > 0 {
			if err = tx.Model(&system.Authority{}).Association("Menus").Delete(&info.Menus); err != nil {
				return errors.Wrap(err, "删除角色菜单失败!")
			}
		} // 删除角色菜单信息
		if err = tx.Model(&system.Casbin{}).Delete(&list).Error; err != nil {
			return errors.Wrap(err, "删除角色权限失败!")
		} // 删除角色权限信息
		if err = tx.Delete(&[]system.UseAuthority{}, "authority_id = ?", info.AuthorityId).Error; err != nil {
			return errors.Wrap(err, "删除用户多角色失败!")
		} // 删除用户多角色信息
		return nil
	})
}

// GetAuthorityInfo 获取角色所有信息
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) GetAuthorityInfo(info *common.GetAuthorityId) (*system.Authority, error) {
	var entity system.Authority
	err := global.Db.Preload("Resources").Where("authority_id = ?", info.AuthorityId).First(&entity).Error
	return &entity, err
}

// SetAuthorityResources 设置角色资源权限
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) SetAuthorityResources(info *request.AuthoritySetResources) error {
	var entity system.Authority
	if err := global.Db.Preload("Resources").First(&entity, "authority_id = ?", info.AuthorityId).Error; err != nil {
		return errors.Wrap(err, "角色查找失败!")
	}
	if err := global.Db.Model(&entity).Association("Resources").Replace(&info.Resources); err != nil {
		return errors.Wrap(err, "设置角色资源权限!")
	}
	return nil
}

// SetAuthorityMenu 菜单与角色绑定
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) SetAuthorityMenu(info *request.AuthoritySetMenu) error {
	var entity system.Authority
	if err := global.Db.Preload("Menus").First(&entity, "authority_id = ?", info.AuthorityId).Error; err != nil {
		return errors.Wrap(err, "菜单查找失败!")
	}
	if err := global.Db.Model(&entity).Association("Menus").Replace(&info.Menus); err != nil {
		return errors.Wrap(err, "设置菜单与角色绑定!")
	}
	return nil
}

// SetAuthorityMenuByTransaction 菜单与角色绑定 事务
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) SetAuthorityMenuByTransaction(tx *gorm.DB, info *request.AuthoritySetMenu) error {
	var entity system.Authority
	if err := tx.Preload("Menus").First(&entity, "authority_id = ?", info.AuthorityId).Error; err != nil {
		return errors.Wrap(err, "菜单查找失败!")
	}
	if err := tx.Model(&entity).Association("Menus").Replace(&info.Menus); err != nil {
		return errors.Wrap(err, "设置菜单与角色绑定!")
	}
	return nil
}

// GetList 分页获取数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *authority) GetList(info *common.PageInfo) (list []system.Authority, total int64, err error) {
	var entities []system.Authority
	db := global.Db.Model(&system.Authority{})
	err = db.Scopes(common.Paginate(info)).Preload("Resources").Where("parent_id = 0").Find(&entities).Error
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
func (s *authority) findChildrenAuthority(authority *system.Authority) error {
	err := global.Db.Preload("Resources").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = s.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
