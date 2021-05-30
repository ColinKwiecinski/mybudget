package auth

import (
	"time"

	"mybudget.com/src/backend/users"
)

type SessionState struct {
	StartTime time.Time   `json:"time"`
	User      *users.User `json:"user"`
}
