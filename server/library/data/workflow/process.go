package data

import (
	model "gf-vue-admin/app/model/extra"
	"gf-vue-admin/library/global"
	"gorm.io/gorm"
	"time"
)

var (
	Process = new(_process)
	processes = []model.WorkflowProcess{{ID: "leaveFlow", CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: "leaveFlow", Clazz: "process", Label: "请假流程（演示）", HideIcon: false, Description: "请假流程演示", View: "view/iconList/index.vue"}}
)

type _process struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: workflow_processes 表数据初始化
func (p *_process) Init() error {
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
