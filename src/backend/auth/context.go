package auth

import (
	"mybudget.com/src/backend/sessions"
	"mybudget.com/src/backend/users"
)

type HandlerContext struct {
	SigningKey string
	Sessions   sessions.Store
	Users      users.Store
}
