package routes

import (
	"github.com/andrres017/technical-test/controllers"
	"github.com/labstack/echo/v4"
)

func ProgramParticipantRoutes(e *echo.Group) {
	e.POST("/program-participants", controllers.CreateProgramParticipantHandler)
	e.GET("/program-participants", controllers.FetchProgramParticipantsHandler)
	e.GET("/program-participants/:id", controllers.GetProgramParticipantByIDHandler)
	e.PUT("/program-participants/:id", controllers.UpdateProgramParticipantHandler)
	e.DELETE("/program-participants/:id", controllers.DeleteProgramParticipantHandler)
}
