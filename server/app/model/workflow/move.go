package model

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/global"
)

type WorkflowMove struct {
	global.Model
	WorkflowProcessID string          `json:"workflowProcessID" gorm:"comment:工作流模板ID"`
	WorkflowProcess   WorkflowProcess `gorm:"<-:false" json:"workflowProcess" gorm:"comment:工作流模板具体信息"`
	WorkflowNodeID    string          `json:"workflowNodeID" gorm:"comment:工作流节点ID"`
	WorkflowNode      WorkflowNode    `gorm:"<-:false" json:"workflowNode" gorm:"comment:工作流节点具体信息"`
	BusinessType      string          `json:"businessType" gorm:"comment:业务标记"`
	BusinessID        uint            `json:"businessID" gorm:"comment:业务ID"`
	PromoterID        uint            `json:"promoterID" gorm:"comment:当前流转发起人"`
	Promoter          model.Admin     `gorm:"<-:false" json:"promoter" gorm:"comment:当前流转发起人信息"`
	OperatorID        uint            `json:"operatorID" gorm:"comment:当前流转操作人"`
	Operator          model.Admin     `gorm:"<-:false" json:"operator" gorm:"comment:当前流转操作人信息"`
	Action            string          `json:"action" gorm:"comment:工作流驱动事件"`
	Param             string          `json:"param" gorm:"comment:工作流驱动参数"`
	IsActive          bool            `json:"isActive" gorm:"comment:是否是活跃节点 "`
}