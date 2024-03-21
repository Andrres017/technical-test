package controllers

import (
	"net/http"
	"strconv"

	"github.com/andrres017/technical-test/models"
	"github.com/andrres017/technical-test/services"
	"github.com/andrres017/technical-test/utils"
	"github.com/labstack/echo/v4"
)

// CreateCompanyHandler handles the request to create a new company.
func CreateCompanyHandler(c echo.Context) error {
	var company models.Companies
	if err := c.Bind(&company); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", "Invalid data", nil)
	}
	if err := c.Validate(company); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", err.Error(), nil)
	}
	createdCompany, err := services.CreateCompany(company)
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Creation Error", "Failed to create company", nil)
	}
	return utils.ApiResponse(c, http.StatusCreated, "Company Created", "The company has been successfully created", createdCompany)
}

// FetchCompaniesHandler handles the request to fetch all companies with pagination.
func FetchCompaniesHandler(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10 // A default value for pageSize
	}
	companies, totalRows, err := services.FetchCompanies(page, pageSize)
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Fetch Error", "Failed to fetch companies", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Companies Fetched", "", echo.Map{
		"data":      companies,
		"totalRows": totalRows,
	})
}

// GetCompanyByIDHandler handles the request to get a specific company by ID.
func GetCompanyByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	company, err := services.GetCompanyByID(uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusNotFound, "Not Found", "Company not found", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Company Found", "", company)
}

// UpdateCompanyHandler handles the request to update an existing company.
func UpdateCompanyHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var company models.Companies
	if err := c.Bind(&company); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", "Invalid data", nil)
	}
	updatedCompany, err := services.UpdateCompany(company, uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Update Error", "Failed to update company", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Company Updated", "The company has been successfully updated", updatedCompany)
}

// DeleteCompanyHandler handles the request to delete a company by its ID.
func DeleteCompanyHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.DeleteCompany(uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Deletion Error", "Failed to delete company", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Company Deleted", "The company has been successfully deleted", nil)
}
