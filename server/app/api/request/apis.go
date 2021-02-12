package request

import (
	model "gf-vue-admin/app/model/system"
)

type BaseApi struct {
	Path        string `p:"path" v:"path@required|length:1,100#请输入api路径|api路径长度为:min到:max位"`
	Method      string `p:"method" v:"method@required|length:1,100#请输入api请求方法|api请求方法长度为:min到:max位"`
	ApiGroup    string `p:"apiGroup" v:"apiGroup@required|length:1,100#请输入api组|api组描述长度为:min到:max位"`
	Description string `p:"description" v:"description@required|length:1,100#请输入api中文描述|api中文描述长度为:min到:max位"`
}

type CreateApi struct {
	BaseApi
}

func (c *CreateApi) Create() *model.Api {
	return &model.Api{
		Path:        c.Path,
		Method:      c.Method,
		ApiGroup:    c.ApiGroup,
		Description: c.Description,
	}
}

type UpdateApi struct {
	GetById
	BaseApi
}

func (u *UpdateApi) Update() *model.Api {
	return &model.Api{
		Path:        u.Path,
		Method:      u.Method,
		ApiGroup:    u.ApiGroup,
		Description: u.Description,
	}
}

// api分页条件查询及排序结构体
type GetApiList struct {
	Path        string `p:"path"`
	Description string `p:"description"`
	ApiGroup    string `p:"apiGroup"`
	Method      string `p:"method"`
	PageInfo
}
