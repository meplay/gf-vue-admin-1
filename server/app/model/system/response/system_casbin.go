package response

import "flipped-aurora/gf-vue-admin/server/app/model/system/request"

type PolicyPath struct {
	Paths []request.CasbinInfo `json:"paths"`
}
