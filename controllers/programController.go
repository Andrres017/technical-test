package controllers

import (
	"net/http"
	"strconv"

	"github.com/andrres017/technical-test/models"
	"github.com/andrres017/technical-test/services"
	"github.com/andrres017/technical-test/utils"
	"github.com/labstack/echo/v4"
)

// CreateProgramHandler handles the request to create a new program.
func CreateProgramHandler(c echo.Context) error {
	var program models.Program
	if err := c.Bind(&program); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", "Invalid data", nil)
	}
	if err := c.Validate(program); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", err.Error(), nil)
	}
	createdProgram, err := services.CreateProgram(program)
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Creation Error", "Failed to create the program", nil)
	}
	return utils.ApiResponse(c, http.StatusCreated, "Program Created", "The program has been successfully created", createdProgram)
}

// FetchProgramsHandler handles the request to fetch all programs with pagination.
func FetchProgramsHandler(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10 // A default value for pageSize
	}
	programs, totalRows, err := services.FetchPrograms(page, pageSize)
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Fetch Error", "Failed to fetch programs", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Programs Fetched", "", echo.Map{"data": programs, "totalRows": totalRows})
}

// GetProgramByIDHandler searches for a specific program by ID.
func GetProgramByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	program, err := services.GetProgramByID(uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusNotFound, "Not Found", "Program not found", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Program Found", "", program)
}

// UpdateProgramHandler handles the request to update an existing program.
func UpdateProgramHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var program models.Program
	if err := c.Bind(&program); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", "Invalid data", nil)
	}
	updatedProgram, err := services.UpdateProgram(program, uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Update Error", "Failed to update the program", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Program Updated", "The program has been successfully updated", updatedProgram)
}

// DeleteProgramHandler handles the request to delete a program by its ID.
func DeleteProgramHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.DeleteProgram(uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Deletion Error", "Failed to delete the program", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Program Deleted", "The program has been successfully deleted", nil)
}
