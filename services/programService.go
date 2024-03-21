package services

import (
	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
)

// CreateProgram crea un nuevo programa en la base de datos.
func CreateProgram(program models.Program) (models.Program, error) {
	result := database.DB.Create(&program)
	return program, result.Error
}

// FetchPrograms recupera programas con paginación de la base de datos.
func FetchPrograms(page int, pageSize int) ([]models.Program, int64, error) {
	var programs []models.Program
	var totalRows int64 = 0
	offset := (page - 1) * pageSize
	result := database.DB.Offset(offset).Limit(pageSize).Find(&programs)
	database.DB.Model(&models.Program{}).Count(&totalRows)
	return programs, totalRows, result.Error
}

// GetProgramByID busca un programa por su ID.
func GetProgramByID(id uint) (models.Program, error) {
	var program models.Program
	result := database.DB.First(&program, id)
	return program, result.Error
}

// UpdateProgram actualiza un programa existente.
func UpdateProgram(program models.Program, id uint) (models.Program, error) {
	if err := database.DB.Model(&models.Program{}).Where("id = ?", id).Updates(program).Error; err != nil {
		return models.Program{}, err
	}
	return program, nil
}

// DeleteProgram elimina un programa por su ID.
func DeleteProgram(id uint) error {
	result := database.DB.Delete(&models.Program{}, id)
	return result.Error
}
