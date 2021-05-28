package auth

import "std/mybudget/src/backend/sessions"

type HandlerContext struct {
	SigningKey string
	Sessions   sessions.Store
	Users      Store
}
