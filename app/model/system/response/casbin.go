package response

import "github.com/flipped-aurora/gf-vue-admin/app/model/system/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
