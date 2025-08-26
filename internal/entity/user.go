package entity

import "time"

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
	UserRoleGuest UserRole = "guest"
)

func (r UserRole) Valid() (ok bool) {
	switch r {
	case UserRoleAdmin, UserRoleUser, UserRoleGuest:
		return true
	default:
		return false
	}
}

type (
	User struct {
		ID        string
		Name      string
		Email     string
		Role      UserRole
		CreatedAt time.Time
	}
)
