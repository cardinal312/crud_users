package models

import (
	"testing"
	"time"
)

// TestUser ...

func TestUser(t *testing.T) *User {
	return &User{
		ID:        1,
		Firstname: "John",
		Lastname:  "Dou",
		Email:     "user@example.org",
		Created:   time.Now(),
	}
}
