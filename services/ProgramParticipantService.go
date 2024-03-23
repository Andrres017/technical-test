package services

import (
	"errors"
	"fmt"

	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
	"gorm.io/gorm"
)

// CreateProgramParticipant crea una nueva asociación de participante de programa en la base de datos.
func CreateProgramParticipant(pp models.ProgramParticipant) (models.ProgramParticipant, error) {
	fmt.Print(pp.ParticipantType)
	// Verifica que ProgramID y ParticipantID sea
	result := database.DB.Create(&pp)
	if result.Error != nil {
		return models.ProgramParticipant{}, result.Error
	}
	return pp, nil
}

// checkParticipantExists verifica la existencia de un participante según el tipo proporcionado.
func checkParticipantExists(id uint, model interface{}) bool {
	result := database.DB.First(model, id)
	return result.Error == nil
}

// FetchProgramParticipants recupera participantes de programas con paginación de la base de datos.
func FetchProgramParticipants(page int, pageSize int) ([]models.ProgramParticipantDetail, int64, error) {
	var programParticipants []models.ProgramParticipant
	var details []models.ProgramParticipantDetail
	var totalRows int64 = 0

	offset := (page - 1) * pageSize

	// Encuentra los participantes del programa y cuenta el total para paginación
	result := database.DB.Offset(offset).Limit(pageSize).Find(&programParticipants)
	database.DB.Model(&models.ProgramParticipant{}).Count(&totalRows)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	// Precarga o agrupa consultas aquí según sea necesario
	// La implementación específica dependerá de tus modelos y estructura de la base de datos
	// Esta es una simplificación y se debe adaptar a tus necesidades

	// Por cada participante del programa, buscar detalles adicionales
	for _, pp := range programParticipants {
		detail := models.ProgramParticipantDetail{
			ProgramParticipant: pp,
		}

		// Buscar el programa asociado
		var program models.Program
		if err := database.DB.First(&program, pp.ProgramID).Error; err == nil {
			detail.Program = program
		} else if err != gorm.ErrRecordNotFound {
			return nil, 0, err // Maneja el error según sea necesario
		}

		// Buscar información detallada del participante basada en el ParticipantType
		switch pp.ParticipantType {
		case models.UserType:
			var user models.User
			if err := database.DB.First(&user, pp.ParticipantID).Error; err == nil {
				detail.ParticipantDetail = user
			}
		case models.ChallengeType:
			var challenge models.Challenge
			if err := database.DB.First(&challenge, pp.ParticipantID).Error; err == nil {
				detail.ParticipantDetail = challenge
			}
		case models.CompanyType:
			var company models.Companies
			if err := database.DB.First(&company, pp.ParticipantID).Error; err == nil {
				detail.ParticipantDetail = company
			}
		}

		details = append(details, detail)
	}

	return details, totalRows, nil
}

// GetProgramParticipantByID busca una asociación de participante de programa por su ID.
func GetProgramParticipantByID(id uint) (models.ProgramParticipantDetail, error) {
	var detail models.ProgramParticipantDetail
	var programParticipant models.ProgramParticipant
	result := database.DB.First(&programParticipant, id)
	if result.Error != nil {
		return detail, result.Error
	}

	// Busca y adjunta el programa asociado
	var program models.Program
	if err := database.DB.First(&program, programParticipant.ProgramID).Error; err != nil {
		return detail, err
	}
	detail.Program = program

	// Adjunta información detallada del participante basada en ParticipantType
	switch programParticipant.ParticipantType {
	case models.UserType:
		var user models.User
		if err := database.DB.First(&user, programParticipant.ParticipantID).Error; err == nil {
			detail.ParticipantDetail = user
		}
	case models.ChallengeType:
		var challenge models.Challenge
		if err := database.DB.First(&challenge, programParticipant.ParticipantID).Error; err == nil {
			detail.ParticipantDetail = challenge
		}
	case models.CompanyType:
		var company models.Companies
		if err := database.DB.First(&company, programParticipant.ParticipantID).Error; err == nil {
			detail.ParticipantDetail = company
		}
	}

	detail.ProgramParticipant = programParticipant
	return detail, nil
}

func UpdateProgramParticipant(pp models.ProgramParticipant, id uint) (models.ProgramParticipant, error) {
	// Verifica primero si la asociación de participante de programa existe.
	var existingPP models.ProgramParticipant
	if err := database.DB.First(&existingPP, id).Error; err != nil {
		return models.ProgramParticipant{}, fmt.Errorf("program participant with ID %d not found", id)
	}

	// Verifica la existencia del programa.
	var program models.Program
	if err := database.DB.First(&program, pp.ProgramID).Error; err != nil {
		return models.ProgramParticipant{}, fmt.Errorf("program with ID %d not found", pp.ProgramID)
	}

	// Verifica la existencia del participante según el ParticipantType.
	switch pp.ParticipantType {
	case models.UserType:
		if !checkParticipantExists(pp.ParticipantID, &models.User{}) {
			return models.ProgramParticipant{}, fmt.Errorf("user with ID %d not found", pp.ParticipantID)
		}
	case models.ChallengeType:
		if !checkParticipantExists(pp.ParticipantID, &models.Challenge{}) {
			return models.ProgramParticipant{}, fmt.Errorf("challenge with ID %d not found", pp.ParticipantID)
		}
	case models.CompanyType:
		if !checkParticipantExists(pp.ParticipantID, &models.Companies{}) {
			return models.ProgramParticipant{}, fmt.Errorf("company with ID %d not found", pp.ParticipantID)
		}
	default:
		return models.ProgramParticipant{}, errors.New("invalid participant type")
	}

	// Actualización de los campos necesarios.
	existingPP.ProgramID = pp.ProgramID
	existingPP.ParticipantID = pp.ParticipantID
	existingPP.ParticipantType = pp.ParticipantType

	// Guarda los cambios en la base de datos.
	if err := database.DB.Save(&existingPP).Error; err != nil {
		return models.ProgramParticipant{}, err
	}
	return existingPP, nil
}

// DeleteProgramParticipant elimina una asociación de participante de programa por su ID.
func DeleteProgramParticipant(id uint) error {
	var pp models.ProgramParticipant
	// Primero, intenta encontrar la asociación por el ID proporcionado.
	result := database.DB.First(&pp, id)
	if result.Error != nil {
		return errors.New("program participant not found")
	}

	// Si el registro existe, procede con la eliminación.
	if result := database.DB.Delete(&pp); result.Error != nil {
		return result.Error
	}

	return nil
}

// CheckProgramExists verifica si un programa existe.
func CheckProgramExists(programID uint) (bool, error) {
	var program models.Program
	result := database.DB.First(&program, programID)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// CheckParticipantExists verifica si un participante existe, basado en el tipo.
func CheckParticipantExists(participantID uint, participantType models.ParticipantType) (bool, error) {
	switch participantType {
	case models.UserType:
		var user models.User
		result := database.DB.First(&user, participantID)
		if result.Error != nil {
			return false, result.Error
		}
	case models.ChallengeType:
		var challenge models.Challenge
		result := database.DB.First(&challenge, participantID)
		if result.Error != nil {
			return false, result.Error
		}
	case models.CompanyType:
		var company models.Companies
		result := database.DB.First(&company, participantID)
		if result.Error != nil {
			return false, result.Error
		}
	default:
		return false, fmt.Errorf("invalid participant type: %s", participantType)
	}
	return true, nil
}
