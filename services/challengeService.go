package services

import (
	"errors"

	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
)

// CreateChallenge crea un nuevo desafío en la base de datos.
func CreateChallenge(challenge models.Challenge) (models.Challenge, error) {
	// Verifica que el campo 'Name' no esté vacío
	if challenge.Name == "" {
		return models.Challenge{}, errors.New("the 'Name' field is required")
	}

	result := database.DB.Create(&challenge)
	return challenge, result.Error
}

// FetchChallenges recupera desafíos con paginación de la base de datos.
func FetchChallenges(page, pageSize int) ([]models.Challenge, int64, error) {
	var challenges []models.Challenge
	var totalRows int64 = 0
	offset := (page - 1) * pageSize
	result := database.DB.Offset(offset).Limit(pageSize).Find(&challenges)
	database.DB.Model(&models.Challenge{}).Count(&totalRows)
	return challenges, totalRows, result.Error
}

// GetChallengeByID busca un desafío por su ID.
func GetChallengeByID(id uint) (models.Challenge, error) {
	var challenge models.Challenge
	result := database.DB.First(&challenge, id)
	return challenge, result.Error
}

// UpdateChallenge actualiza un desafío existente.
func UpdateChallenge(challenge models.Challenge, id uint) (models.Challenge, error) {
	var existingChallenge models.Challenge
	if err := database.DB.First(&existingChallenge, id).Error; err != nil {
		return models.Challenge{}, err
	}
	existingChallenge.Name = challenge.Name
	if err := database.DB.Save(&existingChallenge).Error; err != nil {
		return models.Challenge{}, err
	}
	return existingChallenge, nil
}

// DeleteChallenge elimina un desafío por su ID.
func DeleteChallenge(id uint) error {
	var challenge models.Challenge
	// Primero, intenta encontrar el Challenge por ID para verificar que exista.
	if err := database.DB.First(&challenge, id).Error; err != nil {
		return errors.New("challenge not found")
	}

	// Si el registro existe, procede con la eliminación.
	if err := database.DB.Delete(&challenge).Error; err != nil {
		return err
	}

	return nil
}
