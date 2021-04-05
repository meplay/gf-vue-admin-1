package response

import "gf-vue-admin/app/model/system/request"

type PolicyPath struct {
	Paths []request.CasbinInfo `json:"paths"`
}
