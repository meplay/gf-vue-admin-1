package response

import (
	"server/app/model/admins"
)

// AdminLogin response Structure
type AdminLogin struct {
	User      *admins.AdminHasOneAuthority `json:"user"`
	Token     string                       `json:"token"`
	ExpiresAt int64                        `json:"expiresAt"`
}

// AdminResponse response Structure
type AdminResponse struct {
	Admin *admins.Entity `json:"user"`
}
