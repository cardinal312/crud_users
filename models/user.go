package models

import (
	"time"
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
