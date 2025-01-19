package handlers

import (
	"net/http"
	"time"

	"github.com/cardinal312/crud_users/internal/app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CreateUser ...
func CreateUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Create user
		var user models.User
		currentTime := time.Now()
		currentTime.Format("2006-01-02 15:04:05")
		user.Created = currentTime

		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  "Error",
				Message: "Could not add the user",
			})
		}

		if error := user.Validate(); error != nil {
			return error
		}

		if err := db.Create(&user).Error; err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  "Error",
				Message: "This user registed in db before",
			})
		}
		return c.JSON(http.StatusCreated, &user)
	}
}

// GetAllUsers ...
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

// GetUser ...
func GetUserById(db *gorm.DB) echo.HandlerFunc {
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

// EditUser ...
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

// DeleteUser ...
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
