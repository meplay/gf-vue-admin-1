package model

import (
	"gorm.io/gorm"
	"time"
)

type WorkflowProcess struct {
	ID          string `json:"id" form:"id" gorm:"comment:流程标识;primaryKey;unique;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Name        string         `json:"name" gorm:"comment:流程名称"`
	Category    string         `json:"category" gorm:"comment:分类"`
	Clazz       string         `json:"clazz" gorm:"comment:类型"`
	Label       string         `json:"label" gorm:"comment:流程标题"`
	HideIcon    bool           `json:"hideIcon" gorm:"comment:是否隐藏图标"`
	Description string         `json:"description" gorm:"comment:详细介绍"`
	View        string         `json:"view" gorm:"comment:前端视图文件"`
	Nodes       []WorkflowNode `json:"nodes"` // 流程节点数据
	Edges       []WorkflowEdge `json:"edges"` // 流程链接数据
}
