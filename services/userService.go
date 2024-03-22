package services

import (
	"errors"

	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
)

// CreateUser crea un nuevo usuario en la base de datos.
func CreateUser(user models.User) (models.User, error) {
	// Validación simple para asegurar que el campo 'Name' no esté vacío.
	if user.Name == "" {
		return models.User{}, errors.New("the 'Name' field is required and cannot be empty")
	}

	result := database.DB.Create(&user)
	return user, result.Error
}

// FetchUsers recupera usuarios con paginación de la base de datos.
func GetUsers(page int, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var totalRows int64 = 0
	offset := (page - 1) * pageSize
	result := database.DB.Offset(offset).Limit(pageSize).Find(&users)
	database.DB.Model(&models.User{}).Count(&totalRows) // Obtener el total de registros
	return users, totalRows, result.Error
}

// GetUserByID busca un usuario por su ID.
func GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	return user, result.Error
}

// UpdateUser actualiza un usuario existente.
func UpdateUser(user models.User, id uint) (models.User, error) {
	var existingUser models.User

	// Primero, intenta encontrar el usuario por el ID proporcionado.
	if err := database.DB.First(&existingUser, id).Error; err != nil {
		return models.User{}, err // Retorna un error si el usuario no se encuentra
	}

	// Ahora que se encontró el usuario, actualiza los campos necesarios.
	// Esto asume que tu objeto 'user' tiene los campos que quieres actualizar.
	// Puedes ajustar esta parte para actualizar campos específicos o toda la entidad.
	existingUser.Name = user.Name
	// Actualiza cualquier otro campo necesario...

	// Guarda los cambios en la base de datos.
	if err := database.DB.Save(&existingUser).Error; err != nil {
		return models.User{}, err // Retorna un error si la actualización falla
	}

	return existingUser, nil // Retorna el usuario actualizado
}

// DeleteUser elimina un usuario por su ID.
func DeleteUser(id uint) error {
	var user models.User
	// Primero, intenta encontrar el usuario por el ID proporcionado.
	if err := database.DB.First(&user, id).Error; err != nil {
		// Si el usuario no se encuentra, retorna un error.
		return errors.New("user not found")
	}

	// Si el usuario existe, procede con la eliminación.
	result := database.DB.Delete(&user)
	return result.Error
}
