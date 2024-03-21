package services

import (
	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
)

// CreateProgramParticipant crea una nueva asociación de participante de programa en la base de datos.
func CreateProgramParticipant(pp models.ProgramParticipant) (models.ProgramParticipant, error) {
	result := database.DB.Create(&pp)
	return pp, result.Error
}

// FetchProgramParticipants recupera participantes de programas con paginación de la base de datos.
func FetchProgramParticipants(page int, pageSize int) ([]models.ProgramParticipant, int64, error) {
	var programParticipants []models.ProgramParticipant
	var totalRows int64 = 0
	offset := (page - 1) * pageSize
	result := database.DB.Offset(offset).Limit(pageSize).Find(&programParticipants)
	database.DB.Model(&models.ProgramParticipant{}).Count(&totalRows)
	return programParticipants, totalRows, result.Error
}

// GetProgramParticipantByID busca una asociación de participante de programa por su ID.
func GetProgramParticipantByID(id uint) (models.ProgramParticipant, error) {
	var programParticipant models.ProgramParticipant
	result := database.DB.First(&programParticipant, id)
	return programParticipant, result.Error
}

// UpdateProgramParticipant actualiza una asociación de participante de programa existente.
func UpdateProgramParticipant(pp models.ProgramParticipant, id uint) (models.ProgramParticipant, error) {
	var existingPP models.ProgramParticipant
	if err := database.DB.First(&existingPP, id).Error; err != nil {
		return models.ProgramParticipant{}, err
	}

	// Actualiza campos específicos aquí, por ejemplo:
	existingPP.ParticipantID = pp.ParticipantID
	existingPP.ParticipantType = pp.ParticipantType

	if err := database.DB.Save(&existingPP).Error; err != nil {
		return models.ProgramParticipant{}, err
	}
	return existingPP, nil
}

// DeleteProgramParticipant elimina una asociación de participante de programa por su ID.
func DeleteProgramParticipant(id uint) error {
	result := database.DB.Delete(&models.ProgramParticipant{}, id)
	return result.Error
}
