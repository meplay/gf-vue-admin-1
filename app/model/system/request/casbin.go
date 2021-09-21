package request

import "gorm.io/gorm"

type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// CasbinInReceive structure for input parameters
type CasbinInReceive struct {
	AuthorityId string       `json:"authorityId"` // 权限id
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

type CasbinSearch struct {
	AuthorityId string       `json:"authorityId"` // 权限id
}

func (r *CasbinSearch) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.AuthorityId != "" {
			db = db.Where("v0 = ?", r.AuthorityId)
		}
		return db
	}
}