package data

import (
	model "gf-vue-admin/app/model/workflow"
	system "gf-vue-admin/library/data/system"
	"gf-vue-admin/library/global"
	"time"

	"gorm.io/gorm"
)

var (
	Process = new(_process)
)

type _process struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: workflow_processes 表数据初始化
func (p *_process) Init() error {
	var processes = []model.WorkflowProcess{
		{ID: "leaveFlow", CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: "leaveFlow", Clazz: "process", Label: system.I18nHash["LeaveProcess"], HideIcon: false, Description: system.I18nHash["LeaveProcess"], View: "view/iconList/index.vue"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&processes).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 定义表名
func (p *_process) TableName() string {
	return "workflow_processes"
}
