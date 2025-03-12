package enums

type UserStatus string

const (
	UserActive   UserStatus = "active"
	UserInactive UserStatus = "inactive"
	UserBanned   UserStatus = "banned"
)
