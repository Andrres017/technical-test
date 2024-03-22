package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/models"
	"github.com/andrres017/technical-test/routes"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUserController(t *testing.T) {
	// Configura la base de datos en memoria.
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)
	database.DB = db // Asegúrate de que tu capa de servicio use esta instancia de DB para las operaciones.
	assert.NoError(t, db.AutoMigrate(&models.User{}))

	// Configura Echo y registra las rutas.
	e := echo.New()
	v1 := e.Group("/api/v1")
	e.Validator = &CustomValidator{validator: validator.New()}
	routes.UserRoutes(v1) // Registra tus rutas.

	// Crea la solicitud y el recorder.
	user := models.User{Name: "Test User"}
	reqBody, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Ejecuta el servidor con la solicitud.
	e.ServeHTTP(rec, req)

	// Verifica la respuesta.
	assert.Equal(t, http.StatusCreated, rec.Code)
	var response map[string]interface{}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
	assert.Equal(t, "The user has been successfully created", response["message"])
}

// func TestGetUsersPaginatedController(t *testing.T) {
// 	// Configura la base de datos en memoria.
// 	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
// 	assert.NoError(t, err)
// 	database.DB = db
// 	assert.NoError(t, db.AutoMigrate(&models.User{}))

// 	// Añade algunos usuarios para probar la paginación.
// 	for i := 0; i < 15; i++ {
// 		user := models.User{Name: "Test User " + strconv.Itoa(i)}
// 		database.DB.Create(&user)
// 	}

// 	// Configura Echo y registra las rutas.
// 	e := echo.New()
// 	v1 := e.Group("/api/v1")
// 	e.Validator = &CustomValidator{validator: validator.New()}
// 	routes.UserRoutes(v1) // Asume que tienes una función que registra las rutas de usuario.

// 	// Crea la solicitud y el recorder para la primera página.
// 	req := httptest.NewRequest(http.MethodGet, "/api/v1/users?page=1&pageSize=10", nil)
// 	rec := httptest.NewRecorder()

// 	// Ejecuta el servidor con la solicitud.
// 	e.ServeHTTP(rec, req)

// 	// Verifica la respuesta de la primera página.
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	var responseFirstPage map[string]interface{}
// 	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &responseFirstPage))

// 	usersDataFirstPage, ok := responseFirstPage["data"].([]interface{})
// 	assert.True(t, ok, "Expected 'data' to be a slice of interface{}")
// 	assert.Equal(t, 10, len(usersDataFirstPage)) // Asegúrate de que hay 10 usuarios en la primera página.

// 	// Reinicia el recorder para la segunda solicitud.
// 	rec = httptest.NewRecorder()

// 	// Verifica los usuarios restantes en la segunda página.
// 	req = httptest.NewRequest(http.MethodGet, "/api/v1/users?page=2&pageSize=10", nil)
// 	e.ServeHTTP(rec, req)

// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	var responseSecondPage map[string]interface{}
// 	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &responseSecondPage))

// 	usersDataSecondPage, ok := responseSecondPage["data"].([]interface{})
// 	assert.True(t, ok, "Expected 'data' to be a slice of interface{}")
// 	assert.Equal(t, 5, len(usersDataSecondPage)) // Verifica los usuarios restantes en la segunda página.
// }

// func TestGetUserController(t *testing.T) {
// 	// Configura la base de datos en memoria y auto-migración para el modelo User.
// 	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
// 	assert.NoError(t, err)
// 	database.DB = db
// 	assert.NoError(t, db.AutoMigrate(&models.User{}))

// 	// Inserta un usuario de prueba en la base de datos.
// 	testUser := models.User{Name: "John Doe"}
// 	result := db.Create(&testUser)
// 	assert.NoError(t, result.Error)

// 	// Configura Echo, el validador, y las rutas.
// 	e := echo.New()
// 	e.Validator = &CustomValidator{validator: validator.New()}
// 	v1 := e.Group("/api/v1")
// 	routes.UserRoutes(v1) // Asume que esta función registra la ruta GetUser.

// 	// Caso de prueba 1: Usuario existente
// 	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/"+strconv.Itoa(int(testUser.ID)), nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/users/:id")
// 	c.SetParamNames("id")
// 	c.SetParamValues(strconv.Itoa(int(testUser.ID)))

// 	if assert.NoError(t, cGetUser(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		var response map[string]interface{}
// 		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
// 		assert.Equal(t, "User Found", response["message"])
// 		assert.NotNil(t, response["data"])
// 	}

// 	// Caso de prueba 2: Usuario no existente
// 	req = httptest.NewRequest(http.MethodGet, "/api/v1/users/999999", nil)
// 	rec = httptest.NewRecorder()
// 	c = e.NewContext(req, rec)
// 	c.SetPath("/users/:id")
// 	c.SetParamNames("id")
// 	c.SetParamValues("999999")

// 	if assert.NoError(t, GetUser(c)) {
// 		assert.Equal(t, http.StatusNotFound, rec.Code)
// 		var response map[string]interface{}
// 		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))
// 		assert.Equal(t, "User not found", response["message"])
// 	}
// }

// CustomValidator es un envoltorio para el validador
type CustomValidator struct {
	validator *validator.Validate
}

// Validate implementa la interfaz de validador de Echo
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
