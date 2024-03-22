package services

import "github.com/andrres017/technical-test/models"

// IUserService define la interfaz para el servicio de usuario.
type IUserService interface {
	CreateUser(user models.User) (models.User, error)
	GetUsers(page, pageSize int) ([]models.User, int64, error)
	GetUserByID(id uint) (models.User, error)
	UpdateUser(user models.User, id uint) (models.User, error)
	DeleteUser(id uint) error
}
