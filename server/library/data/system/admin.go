package data

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
	"time"

	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var (
	Admin = new(admin)
)

type admin struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: admins 表数据初始化
func (a *admin) Init() error {
	var admins = []model.Admin{
		{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Uuid: uuid.NewV4().String(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", Nickname: I18nHash["SuperAdmin"], Avatar: "http://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
		{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Uuid: uuid.NewV4().String(), Username: "a303176530", Password: "3ec063004a6f31642261936a379fde3d", Nickname: I18nHash["OtherUser"], Avatar: "http://qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
	}
	_ = admins[0].EncryptedPassword()
	_ = admins[1].EncryptedPassword()

	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.Admin{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> admins 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 定义表名
func (a *admin) TableName() string {
	return "admins"
}
