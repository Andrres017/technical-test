package main

import (
	"fmt"
	"os"

	"github.com/andrres017/technical-test/database"
	"github.com/andrres017/technical-test/routes"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	godotenv.Load()
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Iniciar conexión a la base de datos
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	database.Connect(dsn)

	// Crea un grupo de rutas con el prefijo 'api/v1'
	v1 := e.Group("/api/v1")

	routes.UserRoutes(v1)
	routes.ChallengeRoutes(v1)
	routes.CompaniesRoutes(v1)
	routes.ProgramRoutes(v1)
	routes.GptAutoFillRoutes(v1)
	routes.ProgramParticipantRoutes(v1)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
