package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cardinal312/crud_users/handlers"
	"github.com/cardinal312/crud_users/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupTestDB() {
	dsn := "host=localhost user=postgres password=example dbname=exampledb port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	db.AutoMigrate(&models.User{})
}

func TestCreateUser(t *testing.T) {
	e := echo.New()

	user := models.User{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		Age:       30,
	}

	// Mock the request
	jsonData, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(jsonData))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the handler
	if assert.NoError(t, handlers.CreateUser(db)(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
