package response

import "gf-vue-admin/app/api/request"

type PolicyPath struct {
	Paths []request.CasbinInfo `json:"paths"`
}
