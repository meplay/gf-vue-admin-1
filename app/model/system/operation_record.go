package system

import (
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"time"
)

type OperationRecord struct {
	global.Model
	Ip           string        `json:"ip" gorm:"column:ip;comment:请求ip"`
	Body         string        `json:"body" gorm:"type:longtext;column:body;comment:请求Body"`
	Path         string        `json:"path" gorm:"column:path;comment:请求路径"`
	Method       string        `json:"method" gorm:"column:method;comment:请求方法"`
	Status       int           `json:"status" gorm:"column:status;comment:请求状态"`
	Agent        string        `json:"agent" gorm:"column:agent;comment:代理"`
	Latency      time.Duration `json:"latency" gorm:"column:latency;comment:延迟" swaggertype:"string"`
	Response     string        `json:"resp" gorm:"type:longtext;column:response;comment:响应Body"`
	ErrorMessage string        `json:"error_message" gorm:"column:error_message;comment:错误信息"`
	UserID       int           `json:"user_id" gorm:"column:user_id;comment:用户id"`
	User         User          `json:"user"`
}
