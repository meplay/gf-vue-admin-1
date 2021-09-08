package response

import "github.com/flipped-aurora/gf-vue-admin/app/model/system"

type UserLogin struct {
	User      *system.User `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
}
