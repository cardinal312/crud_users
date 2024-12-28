package main

import (
	"fmt"
	"log"

	"github.com/cardinal312/crud_users/handlers"
	"github.com/cardinal312/crud_users/models"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MARK: - DATA BASE
var db *gorm.DB

func initDB() {
	// PostgreSQL connection string
	dsn := "host=postgres user=user password=qwerty dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	db.AutoMigrate(&models.User{})
}

// MARK: - MAIN
func main() {

	initDB()
	e := echo.New()

	e.POST("/users", handlers.CreateUser(db))
	e.GET("/users", handlers.GetAllUsers(db))
	e.GET("/user/:id", handlers.GetUser(db))
	e.PATCH("/user/:id", handlers.EditUser(db))
	e.DELETE("/user/:id", handlers.DeleteUser(db))

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
	fmt.Println("Server started at :9090...🙌🏻🔥")
}
