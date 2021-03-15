package model

import (
	"gf-vue-admin/library/global"
)

var WorkflowBusinessStruct map[string]func() Workflow
var WorkflowBusinessTable map[string]func() interface{}

type Workflow interface {
	CreateWorkflowMove() *WorkflowMove
	GetBusinessType() string
	GetBusinessID() uint
	GetWorkflowBase() WorkflowBase
}

type WorkflowBase struct {
	WorkflowMoveID    uint   `json:"workflowMoveID" gorm:"-"`
	BusinessID        uint   `json:"businessID" gorm:"<-:false;column:id"` // 业务对应ID（businessID）的返回
	BusinessType      string `json:"businessType" gorm:"-"`
	PromoterID        uint   `json:"promoterID" gorm:"-"`
	OperatorID        uint   `json:"operatorID" gorm:"-"`
	WorkflowProcessID string `json:"workflowProcessID" gorm:"-"`
	WorkflowNodeID    string `json:"workflowNodeID" gorm:"-"`
	Param             string `json:"param" gorm:"-"`
	Action            string `json:"action" gorm:"-"`
}

func (w WorkflowBase) CreateWorkflowMove() (businessModel *WorkflowMove) {
	return &WorkflowMove{
		Model:             global.Model{ID: w.WorkflowMoveID},
		BusinessType:      w.BusinessType,
		PromoterID:        w.PromoterID,
		OperatorID:        w.OperatorID,
		Param:             w.Param,
		WorkflowProcessID: w.WorkflowProcessID,
		WorkflowNodeID:    w.WorkflowNodeID,
		BusinessID:        w.BusinessID,
		Action:            w.Action,
		IsActive:          true,
	}
}

func (w WorkflowBase) GetBusinessType() (businessType string) {
	return w.BusinessType
}

func (w WorkflowBase) GetBusinessID() (businessID uint) {
	return w.BusinessID
}

func (w WorkflowBase) GetWorkflowBase() (workflowBase WorkflowBase) {
	return w
}



