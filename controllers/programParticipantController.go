package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andrres017/technical-test/models"
	"github.com/andrres017/technical-test/services"
	"github.com/andrres017/technical-test/utils"
	"github.com/labstack/echo/v4"
)

// CreateProgramParticipantHandler handles the creation of program participants.
func CreateProgramParticipantHandler(c echo.Context) error {
	var programParticipant models.ProgramParticipant
	if err := c.Bind(&programParticipant); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", "Invalid data", nil)
	}

	// Verificar la existencia del programa.
	exists, err := services.CheckProgramExists(programParticipant.ProgramID)
	if err != nil || !exists {
		return utils.ApiResponse(c, http.StatusNotFound, "Not Found", "The specified program was not found", nil)
	}

	// Verificar la existencia del participante basado en el tipo.
	exists, err = services.CheckParticipantExists(programParticipant.ParticipantID, programParticipant.ParticipantType)
	if err != nil || !exists {
		return utils.ApiResponse(c, http.StatusNotFound, "Not Found", fmt.Sprintf("The specified participant with type %s was not found", programParticipant.ParticipantType), nil)
	}

	// Si todo está en orden, proceder con la creación.
	createdProgramParticipant, err := services.CreateProgramParticipant(programParticipant)
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Creation Error", "Failed to create program participant", nil)
	}
	return utils.ApiResponse(c, http.StatusCreated, "Program Participant Created", "The program participant has been successfully created", createdProgramParticipant)
}

// FetchProgramParticipantsHandler handles the retrieval of program participants with pagination.
func FetchProgramParticipantsHandler(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10 // Default value for pageSize
	}
	programParticipants, totalRows, err := services.FetchProgramParticipants(page, pageSize)
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Fetch Error", "Failed to fetch program participants", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Program Participants Fetched", "", echo.Map{
		"data":      programParticipants,
		"totalRows": totalRows,
	})
}

// GetProgramParticipantByIDHandler searches for a program participant by its ID.
func GetProgramParticipantByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	programParticipant, err := services.GetProgramParticipantByID(uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusNotFound, "Not Found", "Program participant not found", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Program Participant Found", "", programParticipant)
}

// UpdateProgramParticipantHandler updates an existing program participant.
func UpdateProgramParticipantHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", "Invalid ID format", nil)
	}

	var programParticipant models.ProgramParticipant
	if err := c.Bind(&programParticipant); err != nil {
		return utils.ApiResponse(c, http.StatusBadRequest, "Validation Error", "Invalid data", nil)
	}

	// Verificar la existencia del programa y del participante.
	programExists, err := services.CheckProgramExists(programParticipant.ProgramID)
	if err != nil || !programExists {
		return utils.ApiResponse(c, http.StatusNotFound, "Not Found", "The specified program was not found", nil)
	}

	participantExists, err := services.CheckParticipantExists(programParticipant.ParticipantID, programParticipant.ParticipantType)
	if err != nil || !participantExists {
		return utils.ApiResponse(c, http.StatusNotFound, "Not Found", "The specified participant was not found", nil)
	}

	updatedProgramParticipant, err := services.UpdateProgramParticipant(programParticipant, uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Update Error", "Failed to update program participant", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Program Participant Updated", "The program participant has been successfully updated", updatedProgramParticipant)
}

// DeleteProgramParticipantHandler deletes a program participant by its ID.
func DeleteProgramParticipantHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.DeleteProgramParticipant(uint(id))
	if err != nil {
		return utils.ApiResponse(c, http.StatusInternalServerError, "Deletion Error", "Failed to delete program participant", nil)
	}
	return utils.ApiResponse(c, http.StatusOK, "Program Participant Deleted", "The program participant has been successfully deleted", nil)
}
