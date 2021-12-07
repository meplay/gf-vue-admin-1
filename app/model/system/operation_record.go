package system

import (
	"time"

	"github.com/flipped-aurora/gf-vue-admin/library/global"
)

type OperationRecord struct {
	global.Model
	Status int `json:"status" gorm:"column:status;comment:请求状态"`
	UserID int `json:"user_id" gorm:"column:user_id;comment:用户id"`

	Ip           string        `json:"ip" gorm:"column:ip;comment:请求ip"`
	Path         string        `json:"path" gorm:"column:path;comment:请求路径"`
	Method       string        `json:"method" gorm:"column:method;comment:请求方法"`
	Agent        string        `json:"agent" gorm:"column:agent;comment:代理"`
	Latency      time.Duration `json:"latency" gorm:"column:latency;comment:延迟" swaggertype:"string"`
	Request      string        `json:"body" gorm:"type:text;column:request;comment:请求body"`
	Response     string        `json:"resp" gorm:"type:text;column:response;comment:响应Body"`
	ErrorMessage string        `json:"error_message" gorm:"column:error_message;comment:错误信息"`
	User         User          `json:"user"`
}

func (o *OperationRecord) TableName() string {
	return "system_operation_records"
}
