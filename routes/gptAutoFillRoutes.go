package routes

import (
	"github.com/andrres017/technical-test/controllers"
	"github.com/labstack/echo/v4"
)

func GptAutoFillRoutes(e *echo.Group) {
	// Tus otras rutas aquí...

	// Ruta para GPT Auto Fill
	e.GET("/gpt-auto-fill", controllers.HandleGPTAutoFill)
}
