package controllers

import (
	"net/http"
	"strconv"

	"github.com/andrres017/technical-test/models"
	"github.com/andrres017/technical-test/services"
	"github.com/andrres017/technical-test/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CreateUser controller to create a new user
func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", "Invalid data", nil)
	}

	if err := c.Validate(user); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", err.Error(), nil)
	}

	createdUser, err := services.CreateUser(user)
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Creation Error", "Failed to create user", nil)
	}
	return utils.ApiResponse(c, http.StatusCreated, "User Created", "The user has been successfully created", createdUser)
}

// GetUsersPaginated controller for fetching users with pagination
func GetUsersPaginated(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10 // Default value for pageSize
	}
	users, totalRows, err := services.GetUsers(page, pageSize)
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Fetch Error", "Failed to fetch users", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Users Fetched", "", echo.Map{
		"data":      users,
		"totalRows": totalRows,
	})
}

// GetUser controller to get a user by their ID
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.GetUserByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ApiResponse(c, http.StatusNotFound, "Not Found", "User not found", nil)
		}
		return utils.ApiResponse(c, http.StatusInternalServerError, "Fetch Error", "Failed to get user", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "User Found", "", user)
}

// UpdateUser controller to update an existing user
func UpdateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", "Invalid data", nil)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	updatedUser, err := services.UpdateUser(user, uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Update Error", "Failed to update user", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "User Updated", "The user has been successfully updated", updatedUser)
}

// DeleteUser controller to delete a user by their ID
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.DeleteUser(uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Deletion Error", "Failed to delete user", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "User Deleted", "The user has been successfully deleted", nil)
}
