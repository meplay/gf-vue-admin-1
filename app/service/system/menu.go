package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
)

var Menu = new(_menu)

type _menu struct{}

// Create 添加菜单
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) Create(info *request.MenuCreate) error {
	if !errors.Is(global.Db.Where("name = ?", info.Name).First(&system.Menu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name!")
	}
	return global.Db.Create(&info.Menu).Error
}

// GetList 获取全部路由
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) GetList() (list []system.Menu, total int64, err error) {
	treeMap, _err := s.getTreeMap()
	if _err != nil {
		return nil, 0, _err
	}
	list = treeMap["0"]
	for i := 0; i < len(list); i++ {
		if err = s.getChildrenList(&list[i], treeMap); err != nil {
			return list, total, errors.Wrap(err, "查找子菜单数据失败!")
		}
	}
	return list, total, err
}

// GetTree 获取动态菜单树
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) GetTree() (list []system.Menu, err error) {
	treeMap, _err := s.getTreeMap()
	if _err != nil {
		return nil, _err
	}
	list = treeMap["0"]
	for i := 0; i < len(list); i++ {
		if err = s.getChildrenList(&list[i], treeMap); err != nil {
			return list, errors.Wrap(err, "查找子菜单数据失败!")
		}
	}
	return list, err
}

// AddMenuAuthority 为角色增加menu树
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) AddMenuAuthority(menus []system.Menu, authorityId string) (err error) {
	var entity system.Authority
	entity.AuthorityId = authorityId
	entity.Menus = menus
	err = Authority.SetAuthorityMenu(&request.AuthoritySetMenu{Authority: entity})
	return err
}

// getTreeMap 获取路由总树map
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) getTreeMap() (treeMap map[string][]system.Menu, err error) {
	var entities []system.Menu
	if err = global.Db.Order("sort").Preload("Parameters").Find(&entities).Error; err != nil {
		return nil, errors.Wrap(err, "查找菜单数据失败!")
	}
	treeMap = make(map[string][]system.Menu, len(entities))
	for _, v := range entities {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, nil
}

// getChildrenList 获取菜单的子菜单
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *_menu) getChildrenList(menu *system.Menu, treeMap map[string][]system.Menu) error {
	if value, ok := treeMap[strconv.Itoa(int(menu.ID))]; ok {
		menu.Children = value
		for i := 0; i < len(menu.Children); i++ {
			if err := s.getChildrenList(&menu.Children[i], treeMap); err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("子菜单列表数据获取失败!")
}
