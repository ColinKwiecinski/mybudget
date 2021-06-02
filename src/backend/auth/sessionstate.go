package auth

import (
	"time"
)

type SessionState struct {
	StartTime time.Time `json:"time"`
	User      *User     `json:"user"`
}
