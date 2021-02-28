package request

import "github.com/gogf/gf/frame/g"

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: Paging common input parameter structure
type PageInfo struct {
	Page     int `p:"page" v:"required|length:1,1000#请输入页数|页数长度为:min到:max位" json:"page" form:"page"`
	PageSize int `p:"pageSize" v:"required|length:1,1000#请输入每页大小|每页大小为:min到:max位" json:"pageSize" form:"pageSize"`
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: Get by id structure
type GetById struct {
	Id uint `p:"id" v:"required|length:1,1000#请输入id|id长度为:min到:max位"`
}

func (i *GetById) Condition() g.Map {
	return g.Map{"id": i.Id}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: Get by ids structure
type GetByIds struct {
	Ids []int `p:"ids" v:"required|length:1,1000#请输入ids|ids长度为:min到:max位"`
}

func (i *GetByIds) Condition() g.Map {
	return g.Map{"id IN (?)": g.Slice{i.Ids}}
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: Delete by id structure
type DeleteByIds struct {
	Ids []int `p:"ids" v:"required|length:1,1000#请输入id|id长度为:min到:max位"`
}

type GetByUuid struct {
	Uuid string `p:"uuid"`
}

type GetAuthorityId struct {
	AuthorityId string `p:"authorityId"`
}

func (i *GetAuthorityId) Condition() g.Map {
	return g.Map{"authority_id": i.AuthorityId}
}

func (p *PageInfo) Paginate() (limit, offset int) {
	limit = p.PageSize
	offset = p.PageSize * (p.Page - 1)
	return limit, offset
}
