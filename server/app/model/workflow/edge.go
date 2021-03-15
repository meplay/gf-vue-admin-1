package model

import (
	"gorm.io/gorm"
	"time"
)

type WorkflowEdge struct {
	ID                  string `json:"id" form:"id" gorm:"comment:唯一标识;primaryKey;unique;not null"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt     `json:"-" gorm:"index"`
	WorkflowProcessID   string             `json:"-" gorm:"comment:流程标识"`
	Clazz               string             `json:"clazz" gorm:"comment:类型（线）"`
	Source              string             `json:"source" gorm:"comment:起点节点"`
	Target              string             `json:"target" gorm:"comment:目标节点"`
	SourceAnchor        int                `json:"sourceAnchor" gorm:"comment:起点"`
	TargetAnchor        int                `json:"targetAnchor" gorm:"comment:目标点"`
	Description         string             `json:"description" gorm:"comment:详细介绍"`
	Shape               string             `json:"shape" gorm:"comment:形状"`
	StartPoint          WorkflowStartPoint `json:"startPoint"` // 起点信息
	EndPoint            WorkflowEndPoint   `json:"endPoint"`   // 终点信息
	Label               string             `json:"label" gorm:"comment:标题"`
	HideIcon            bool               `json:"hideIcon" gorm:"comment:隐藏图标"`
	ConditionExpression string             `json:"conditionExpression" gorm:"comment:条件标识"`
	Seq                 string             `json:"seq" gorm:"comment:序号"`
	Reverse             bool               `json:"reverse" gorm:"comment:是否反向"`
}
