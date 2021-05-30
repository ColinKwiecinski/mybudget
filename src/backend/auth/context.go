package auth

import (
	"mybudget.com/src/backend/users"
	"std/mybudget/src/backend/sessions"
)

type HandlerContext struct {
	SigningKey string
	Sessions   sessions.Store
	Users      users.Store
}
