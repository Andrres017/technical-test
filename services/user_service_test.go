package services

import (
	"testing"

	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabaseForUsers(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	database.DB = db
	return db
}

func TestCreateUser(t *testing.T) {
	db := setupDatabaseForUsers(t)
	defer db.Migrator().DropTable(&models.User{})

	user := models.User{Name: "New User"}

	result, err := CreateUser(user)

	assert.NoError(t, err)
	assert.NotZero(t, result.ID)
	assert.Equal(t, "New User", result.Name)
}

func TestGetUsers(t *testing.T) {
	db := setupDatabaseForUsers(t)
	defer db.Migrator().DropTable(&models.User{})

	_, _ = CreateUser(models.User{Name: "User 1"})
	_, _ = CreateUser(models.User{Name: "User 2"})

	users, total, err := GetUsers(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, users, 2)
}

func TestGetUserByID(t *testing.T) {
	db := setupDatabaseForUsers(t)
	defer db.Migrator().DropTable(&models.User{})

	newUser, _ := CreateUser(models.User{Name: "User By ID"})

	foundUser, err := GetUserByID(newUser.ID)

	assert.NoError(t, err)
	assert.Equal(t, "User By ID", foundUser.Name)
}

func TestUpdateUser(t *testing.T) {
	db := setupDatabaseForUsers(t)
	defer db.Migrator().DropTable(&models.User{})

	user, _ := CreateUser(models.User{Name: "Initial Name"})

	user.Name = "Updated Name"
	updatedUser, err := UpdateUser(user, user.ID)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", updatedUser.Name)
}

func TestDeleteUser(t *testing.T) {
	db := setupDatabaseForUsers(t)
	defer db.Migrator().DropTable(&models.User{})

	user, _ := CreateUser(models.User{Name: "To Be Deleted"})

	err := DeleteUser(user.ID)

	assert.NoError(t, err)

	_, err = GetUserByID(user.ID)
	assert.Error(t, err) // Expect an error because the user should no longer exist
}

func TestCreateUserWithError(t *testing.T) {
	db := setupDatabaseForUsers(t)
	defer db.Migrator().DropTable(&models.User{})

	// Intenta crear un User sin nombre, asumiendo que el campo 'Name' es requerido.
	user := models.User{}

	_, err := CreateUser(user)

	// Se espera un error porque el campo 'Name' es requerido y está vacío.
	assert.Error(t, err)
}

func TestGetUserByIDWithError(t *testing.T) {
	db := setupDatabaseForUsers(t)
	defer db.Migrator().DropTable(&models.User{})

	// Intenta buscar un User con un ID que no existe.
	_, err := GetUserByID(999999)

	// Se espera un error porque el ID no existe en la base de datos.
	assert.Error(t, err)
}

func TestUpdateUserWithError(t *testing.T) {
	db := setupDatabaseForUsers(t)
	defer db.Migrator().DropTable(&models.User{})

	// Intenta actualizar un User inexistente.
	user := models.User{ID: 999999, Name: "Nonexistent User"}

	_, err := UpdateUser(user, user.ID)

	// Se espera un error porque el User con el ID especificado no existe.
	assert.Error(t, err)
}

func TestDeleteUserWithError(t *testing.T) {
	db := setupDatabaseForUsers(t)
	defer db.Migrator().DropTable(&models.User{})

	// Intenta eliminar un User que no existe.
	err := DeleteUser(999999)

	// Se espera un error porque el User con el ID especificado no existe.
	assert.Error(t, err)
}
