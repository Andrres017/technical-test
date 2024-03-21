package routes

import (
	"github.com/andrres017/technical-test/controllers"

	"github.com/labstack/echo/v4"
)

func ProgramRoutes(e *echo.Group) {
	e.POST("/programs", controllers.CreateProgramHandler)
	e.GET("/programs", controllers.FetchProgramsHandler)
	e.GET("/programs/:id", controllers.GetProgramByIDHandler)
	e.PUT("/programs/:id", controllers.UpdateProgramHandler)
	e.DELETE("/programs/:id", controllers.DeleteProgramHandler)
}
