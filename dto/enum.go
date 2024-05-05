package dto

type key string

const (
	UserContextKey key = "user"
	AuthCookieName key = "auth"
)

const (
	RequestNew      key = "NEW"
	RequestRejected key = "REJECTED"
	RequestApproved key = "APPROVED"
)

const (
	RoleAdmin key = "ADMIN"
	RoleUser  key = "USER"
)
