package services

import (
	"errors"

	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
)

// CreateProgram crea un nuevo programa en la base de datos.
func CreateProgram(program models.Program) (models.Program, error) {
	// Verifica si el campo 'Name' está vacío y retorna un error si lo está.
	if program.Name == "" {
		return models.Program{}, errors.New("the 'Name' field is required and cannot be empty")
	}

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
	// Primero, verifica si existe el programa que se intenta actualizar.
	var existingProgram models.Program
	if err := database.DB.First(&existingProgram, id).Error; err != nil {
		return models.Program{}, errors.New("program not found")
	}

	// Si el registro existe, procede con la actualización.
	if err := database.DB.Model(&existingProgram).Updates(program).Error; err != nil {
		return models.Program{}, err
	}
	return existingProgram, nil // Retorna el programa actualizado
}

// DeleteProgram elimina un programa por su ID.
func DeleteProgram(id uint) error {
	// Primero, verifica si existe el programa que se intenta eliminar.
	var program models.Program
	if err := database.DB.First(&program, id).Error; err != nil {
		return errors.New("program not found")
	}

	// Si el registro existe, procede con la eliminación.
	result := database.DB.Delete(&program)
	return result.Error
}
