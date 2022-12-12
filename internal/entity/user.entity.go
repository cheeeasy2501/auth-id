package entity

import (
	"time"

	"gorm.io/gorm"

	"github.com/cheeeasy2501/auth-id/internal/transport/http/v1/request"
)

type User struct {
	gorm.Model
	Avatar,
	FirstName,
	LastName,
	MiddleName,
	Password,
	Email string
	emailVerifiedAt time.Time
	isBanned        bool
}

func NewUser() User {
	return User{}
}

func NewUserFromRegistrationRequest(request request.RegistrationRequest) User {
	return User{
		Avatar:     *request.Avatar,
		Email:      request.Email,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		MiddleName: *request.MiddleName,
		Password:   request.Password,
		isBanned:   false,
	}
}

// func (u *User) Password() string {
// 	return u.password
// }

// func (u *User) SetPassword(password string) {
// 	u.password = password
// }
