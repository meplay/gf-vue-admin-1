package request

type CreateCasbinRule struct {
    PType string `p:"pType" v:"required|length:1,1000#请输入pType字段|pType字段长度为:min到:max位"`
    V0 string `p:"v0" v:"required|length:1,1000#请输入v0字段|v0字段长度为:min到:max位"`
    V1 string `p:"v1" v:"required|length:1,1000#请输入v1字段|v1字段长度为:min到:max位"`
    V2 string `p:"v2" v:"required|length:1,1000#请输入v2字段|v2字段长度为:min到:max位"`
    V3 string `p:"v3" v:"required|length:1,1000#请输入v3字段|v3字段长度为:min到:max位"`
    V4 string `p:"v4" v:"required|length:1,1000#请输入v4字段|v4字段长度为:min到:max位"`
    V5 string `p:"v5" v:"required|length:1,1000#请输入v5字段|v5字段长度为:min到:max位"`
}

type UpdateCasbinRule struct {
    PType string `p:"pType" v:"required|length:1,1000#请输入pType字段|pType字段长度为:min到:max位"`
    V0 string `p:"v0" v:"required|length:1,1000#请输入v0字段|v0字段长度为:min到:max位"`
    V1 string `p:"v1" v:"required|length:1,1000#请输入v1字段|v1字段长度为:min到:max位"`
    V2 string `p:"v2" v:"required|length:1,1000#请输入v2字段|v2字段长度为:min到:max位"`
    V3 string `p:"v3" v:"required|length:1,1000#请输入v3字段|v3字段长度为:min到:max位"`
    V4 string `p:"v4" v:"required|length:1,1000#请输入v4字段|v4字段长度为:min到:max位"`
    V5 string `p:"v5" v:"required|length:1,1000#请输入v5字段|v5字段长度为:min到:max位"`
}

type FindCasbinRule struct {
    PType string `p:"pType" v:"required|length:1,1000#请输入pType字段|pType字段长度为:min到:max位"`
    V0 string `p:"v0" v:"required|length:1,1000#请输入v0字段|v0字段长度为:min到:max位"`
    V1 string `p:"v1" v:"required|length:1,1000#请输入v1字段|v1字段长度为:min到:max位"`
    V2 string `p:"v2" v:"required|length:1,1000#请输入v2字段|v2字段长度为:min到:max位"`
    V3 string `p:"v3" v:"required|length:1,1000#请输入v3字段|v3字段长度为:min到:max位"`
    V4 string `p:"v4" v:"required|length:1,1000#请输入v4字段|v4字段长度为:min到:max位"`
    V5 string `p:"v5" v:"required|length:1,1000#请输入v5字段|v5字段长度为:min到:max位"`
}

type GetCasbinRuleList struct {
    PType string `p:"pType"`
    V0 string `p:"v0"`
    V1 string `p:"v1"`
    V2 string `p:"v2"`
    V3 string `p:"v3"`
    V4 string `p:"v4"`
    V5 string `p:"v5"`
	PageInfo
}