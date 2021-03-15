package model

type Base struct {
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
