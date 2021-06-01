package request

import (
	model "gf-vue-admin/app/model/system"
	"gorm.io/gorm"
)

type BaseOperationRecord struct {
	Ip           string `p:"ip"            v:"required | length:1, 20#请输入ip|ip长度为:min到max"`
	Path         string `p:"path"          v:"required | length:1, 20#请输入请求路由|请求路由长度为:min到max"`
	Agent        string `p:"agent"         v:"required | length:1, 20#请输入代理|代理长度为:min到max"`
	Method       string `p:"method"        v:"required | length:1, 20#请输入请求方法|请求方法长度为:min到max"`
	Request      string `p:"request"       v:"required | length:1, 20#请输入请求Body|请求Body长度为:min到max"`
	Response     string `p:"response"      v:"required | length:1, 20#请输入响应Body|响应Body长度为:min到max"`
	ErrorMessage string `p:"error_message" v:"required | length:1, 20#请输入报错信息|报错信息长度为:min到max"`

	Status  int   `p:"status"        v:"required | length:1, 20#请输入状态|状态长度为:min到max"`
	UserID  int   `p:"user_id"       v:"required | length:1, 20#请输入用户id|用户id长度为:min到max"`
	Latency int64 `p:"latency"       v:"required | length:1, 20#请输入延迟|延迟长度为:min到max"`
}

type CreateOperationRecord struct {
	BaseOperationRecord
}

func (c *CreateOperationRecord) Create() *model.OperationRecord {
	return &model.OperationRecord{
		Ip:           c.Ip,
		Path:         c.Path,
		Agent:        c.Agent,
		Method:       c.Method,
		Request:      c.Request,
		Response:     c.Response,
		ErrorMessage: c.Ip,
		Status:       c.Status,
		UserID:       c.UserID,
		Latency:      c.Latency,
	}
}

type SearchOperationRecord struct {
	Path   string `p:"path"`
	Method string `p:"method"`
	Status int    `p:"status"`
	PageInfo
}

func (s *SearchOperationRecord) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if s.Path != "" {
			db.Where("path = ?", s.Path)
		}
		if s.Method != "" {
			db.Where("Method = ?", s.Method)
		}
		if s.Status != 0 {
			db.Where("status = ?", s.Status)
		}
		return db
	}
}
