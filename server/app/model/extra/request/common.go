package request

import "github.com/gogf/gf/frame/g"

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `p:"page" v:"required|length:1,1000#请输入页数|页数长度为:min到:max位" json:"page" form:"page"`
	PageSize int `p:"pageSize" v:"required|length:1,1000#请输入每页大小|每页大小为:min到:max位" json:"pageSize" form:"pageSize"`
}

// GetById Get by id structure
type GetById struct {
	Id uint `p:"id" v:"required|length:1,1000#请输入id|id长度为:min到:max位"`
}

func (i *GetById) Condition() g.Map {
	return g.Map{"id": i.Id}
}

// GetByIds Get by ids structure
type GetByIds struct {
	Ids []int `p:"ids" v:"required|length:1,1000#请输入ids|ids长度为:min到:max位"`
}

func (i *GetByIds) Condition() g.Map {
	return g.Map{"id IN (?)": g.Slice{i.Ids}}
}

// DeleteByIds Delete by id structure
type DeleteByIds struct {
	Ids []int `p:"ids" v:"required|length:1,1000#请输入id|id长度为:min到:max位"`
}

func (p *PageInfo) Paginate() (limit, offset int) {
	limit = p.PageSize
	offset = p.PageSize * (p.Page - 1)
	return limit, offset
}

type ExcelExport struct {
	Filepath  string          `json:"filepath"`
	SheetName string          `json:"sheetName"`
	A1Data    []string        `json:"a1Data"`
	Data      [][]interface{} `json:"data"`
}
