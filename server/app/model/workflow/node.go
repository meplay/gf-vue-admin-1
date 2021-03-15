package model

import (
	"gorm.io/gorm"
	"time"
)

type WorkflowNode struct {
	ID                string `json:"id" form:"id" gorm:"comment:节点id;primaryKey;unique;not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
	WorkflowProcessID string         `json:"workflowProcessID" gorm:"comment:流程标识"`
	Clazz             string         `json:"clazz" gorm:"comment:节点类型"`
	Label             string         `json:"label" gorm:"comment:节点名称"`
	Type              string         `json:"type" gorm:"comment:图标类型"`
	Shape             string         `json:"shape" gorm:"comment:形状"`
	Description       string         `json:"description" gorm:"comment:详细介绍"`
	View              string         `json:"view" gorm:"comment:前端视图文件"`
	X                 float64        `json:"y" gorm:"comment:x位置"`
	Y                 float64        `json:"x" gorm:"comment:y位置"`
	WaitState         string         `json:"waitState" gorm:"comment:等待属性"`
	StateValue        string         `json:"stateValue" gorm:"comment:等待值"`
	To                string         `json:"to" gorm:"comment:收件人"`
	Subject           string         `json:"subject" gorm:"comment:标题"`
	Content           string         `json:"content" gorm:"comment:内容"`
	Cycle             string         `json:"cycle" gorm:"comment:循环时间"`
	Duration          string         `json:"duration" gorm:"comment:持续时间"`
	HideIcon          bool           `json:"hideIcon" gorm:"comment:是否隐藏图标"`
	DueDate           *time.Time     `json:"dueDate" gorm:"comment:到期时间"`
	AssignType        string         `json:"assignType" gorm:"comment:审批类型"`
	AssignValue       string         `json:"assignValue" gorm:"comment:审批类型值"`
	Success           bool           `json:"success" gorm:"comment:是否成功"`
}
