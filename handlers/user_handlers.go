package handlers

import (
	"net/http"
	"time"

	"github.com/cardinal312/crud_users/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// MARK: - Create User
func CreateUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Create user
		var user models.User
		user.Created = time.Now()

		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  "Error",
				Message: "Could not add the user",
			})
		}

		if err := db.Create(&user).Error; err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  "Error",
				Message: "Could not create the user",
			})
		}
		return c.JSON(http.StatusCreated, &user)
	}
}

// MARK: - Get All Users
func GetAllUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var users []models.User

		if err := db.Find(&users).Error; err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  "Error",
				Message: "Could not find the users",
			})
		}
		return c.JSON(http.StatusOK, &users)
	}
}

// MARK: - Get User by ID
func GetUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		var user models.User
		if err := db.First(&user, "id = ?", id).Error; err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  "Error",
				Message: "Could not find the users",
			})
		}
		return c.JSON(http.StatusOK, user)
	}
}

// MARK: - Edit User by ID
func EditUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		var user models.User
		if err := db.First(&user, "id = ?", id).Error; err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  "Error",
				Message: "User not found",
			})
		}

		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  "Error",
				Message: "Invalid input",
			})
		}

		if err := db.Save(&user).Error; err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  "Error",
				Message: "Failed to update user",
			})
		}
		return c.JSON(http.StatusOK, user)
	}
}

// MARK: - Delete User by ID
func DeleteUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		id := c.Param("id")
		var user models.User

		if err := db.First(&user, "id = ?", id).Error; err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  "Error",
				Message: "User not found",
			})
		}

		if err := db.Delete(&user).Error; err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  "Error",
				Message: "Failed to delete user",
			})
		}

		return c.JSON(http.StatusOK, models.Response{
			Status:  "Success",
			Message: "User successfully deleted ",
		})
	}
}
