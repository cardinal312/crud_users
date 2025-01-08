package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	ID        int       `json:"id"`
	Firstname string    `json:"firstname""`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	Created   time.Time `json:"created" gorm:"autoCreateTime"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Validate 
func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.ID, validation.Required),
		validation.Field(&u.Firstname, validation.Required, validation.Length(1, 50)),
		validation.Field(&u.Lastname, validation.Required, validation.Length(1, 50)),
		validation.Field(&u.Email, validation.Required, is.Email))
}
