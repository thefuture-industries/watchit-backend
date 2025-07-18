package types

import "time"

type AuthToken struct {
	UUID      string    `json:"uuid"`
	RoleID    uint      `json:"role_id"`
	Timestamp time.Time `json:"timestamp"`
}
