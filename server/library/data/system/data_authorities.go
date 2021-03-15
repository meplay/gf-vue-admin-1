package data

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var (
	ResourcesAuthorities = new(resources)
	infos                = []model.DataAuthorities{
		{AuthorityId: "888", DataAuthority: "888"},
		{AuthorityId: "888", DataAuthority: "8881"},
		{AuthorityId: "888", DataAuthority: "9528"},
		{AuthorityId: "9528", DataAuthority: "8881"},
		{AuthorityId: "9528", DataAuthority: "9528"},
	}
)

type resources struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: data_authorities 表数据初始化
func (d *resources) Init() error {
	return global.Db.Table("data_authorities").Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ('888', '9528') ").Find(&[]model.DataAuthorities{}).RowsAffected == 5 {
			color.Danger.Println("\n[Mysql] --> data_authorities 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&infos).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> data_authorities 表初始数据成功!")
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 自定义表名
func (d *resources) TableName() string {
	return "data_authorities"
}
