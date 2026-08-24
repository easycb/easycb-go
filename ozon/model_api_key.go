package ozon

import "time"

type GetRolesRsp struct {
	BaseRsp
	ExpiresAt time.Time `json:"expires_at"`
	Roles     []struct {
		Name    string   `json:"name"`
		Methods []string `json:"methods"`
	} `json:"roles,omitempty"`
}
