package controllers

import (
	"net/http"
	"strconv"

	"github.com/andrres017/technical-test/models"
	"github.com/andrres017/technical-test/services"
	"github.com/andrres017/technical-test/utils"
	"github.com/labstack/echo/v4"
)

// CreateChallengeHandler handles the creation of challenges.
func CreateChallengeHandler(c echo.Context) error {
	var challenge models.Challenge
	if err := c.Bind(&challenge); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", "Invalid data", nil)
	}
	if err := c.Validate(&challenge); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", err.Error(), nil)
	}
	createdChallenge, err := services.CreateChallenge(challenge)
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Creation Error", "Failed to create challenge", nil)
	}
	return utils.ApiResponse(c, http.StatusCreated, "Challenge Created", "The challenge has been successfully created", createdChallenge)
}

// FetchChallengesHandler handles the retrieval of challenges with pagination.
func FetchChallengesHandler(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10 // Default value for pageSize
	}
	challenges, totalRows, err := services.FetchChallenges(page, pageSize)
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Fetch Error", "Failed to fetch challenges", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Challenges Fetched", "", echo.Map{
		"data":      challenges,
		"totalRows": totalRows,
	})
}

// GetChallengeByIDHandler searches for a challenge by its ID.
func GetChallengeByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	challenge, err := services.GetChallengeByID(uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusNotFound, "Not Found", "Challenge not found", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Challenge Found", "", challenge)
}

// UpdateChallengeHandler updates an existing challenge.
func UpdateChallengeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var challenge models.Challenge
	if err := c.Bind(&challenge); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", "Invalid data", nil)
	}
	updatedChallenge, err := services.UpdateChallenge(challenge, uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Update Error", "Failed to update challenge", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Challenge Updated", "The challenge has been successfully updated", updatedChallenge)
}

// DeleteChallengeHandler deletes a challenge by its ID.
func DeleteChallengeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.DeleteChallenge(uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Deletion Error", "Failed to delete challenge", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Challenge Deleted", "The challenge has been successfully deleted", nil)
}
