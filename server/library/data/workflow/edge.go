package data

import (
	model "flipped-aurora/gf-vue-admin/server/app/model/workflow"
	system "flipped-aurora/gf-vue-admin/server/library/data/system"
	"flipped-aurora/gf-vue-admin/server/library/global"
	"time"

	"gorm.io/gorm"
)

var Edge = new(edge)

type edge struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: workflow_edges 表数据初始化
func (e *edge) Init() error {
	var edges = []model.WorkflowEdge{
		{ID: "flow1604985849039", CreatedAt: time.Now(), UpdatedAt: time.Now(), WorkflowProcessID: "leaveFlow", Clazz: "flow", Source: "start1603681292875", Target: "userTask1603681299962", SourceAnchor: 1, TargetAnchor: 3, Shape: "flow-polyline-round", Label: "", HideIcon: false, ConditionExpression: "", Reverse: false},
		{ID: "flow1604985879574", CreatedAt: time.Now(), UpdatedAt: time.Now(), WorkflowProcessID: "leaveFlow", Clazz: "flow", Source: "userTask1603681299962", Target: "end1603681360882", SourceAnchor: 0, TargetAnchor: 2, Shape: "flow-polyline-round", Label: system.I18nHash["Agree"], HideIcon: false, ConditionExpression: "yes", Reverse: false},
		{ID: "flow1604985881207", CreatedAt: time.Now(), UpdatedAt: time.Now(), WorkflowProcessID: "leaveFlow", Clazz: "flow", Source: "userTask1603681299962", Target: "end1603681358043", SourceAnchor: 2, TargetAnchor: 2, Shape: "flow-polyline-round", Label: system.I18nHash["Disagree"], HideIcon: false, ConditionExpression: "no", Reverse: false},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&edges).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: 定义表名
func (e *edge) TableName() string {
	return "workflow_edges"
}
