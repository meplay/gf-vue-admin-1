package request

// Casbin info structure
type CasbinInfo struct {
	Path   string `p:"path" json:"path"`
	Method string `p:"method" json:"method"`
}

// Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId string       `p:"authorityId" json:"authorityId"`
	CasbinInfos []CasbinInfo `p:"casbinInfos" json:"casbinInfos"`
}
