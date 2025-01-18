package models_test

import (
	"testing"

	"github.com/cardinal312/crud_users/models"
	"github.com/stretchr/testify/assert"
)

// Test for any invalid error cases
func TestUser_Validate(t *testing.T) {

	testCases := []struct {
		name    string
		u       func() *models.User
		isValid bool
	}{
		/*
			ID:        1,
			Firstname: "John",
			Lastname:  "Dou",
			Email:     "user@example.org",
			Created:   time.Now(),
		*/

		{
			name: "valid",
			u: func() *models.User {
				return models.TestUser(t)
			},
			isValid: true,
		},

		{
			name: "empty email",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},

		{
			name: "invalid email",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},

		{
			name: "first name is not empty",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Firstname = ""
				return u
			},
			isValid: false,
		},

		{
			name: "first name is valid",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Firstname = "John"
				return u
			},
			isValid: true,
		},

		{
			name: "first name max len",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Firstname = "JohnJohnJohnJohnJohnJohnJohnJohnJohnJohnJohnnJohnJohnJohnJohnJohnJohnJohnJohnJohnJohnJohnJohnJohn" // max len(50)
				return u
			},
			isValid: false,
		},

		{
			name: "first name min len",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Firstname = "" // min len(1)
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
