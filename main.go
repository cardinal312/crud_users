package main

import (
	"fmt"
	"log"

	"github.com/cardinal312/crud_users/internal/app/handlers"
	"github.com/cardinal312/crud_users/internal/app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DATA BASE ...
var db *gorm.DB

func initDB() {
	dsn := "host=postgres user=user password=qwerty dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	db.AutoMigrate(&models.User{})
}

// MARK ...
func main() {

	initDB()
	e := echo.New()

	e.POST("/users", handlers.CreateUser(db))
	e.GET("/users", handlers.GetAllUsers(db))
	e.GET("/user/:id", handlers.GetUserById(db))
	e.PATCH("/user/:id", handlers.EditUser(db))
	e.DELETE("/user/:id", handlers.DeleteUser(db))

	// Start server ...
	e.Logger.Fatal(e.Start(":9090"))
	fmt.Println("Server started at localhost:9090...🙌🏻🔥")
}
