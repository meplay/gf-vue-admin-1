package response

import "server/app/api/request"

type PolicyPath struct {
	Paths []request.CasbinInfo `json:"paths"`
}
