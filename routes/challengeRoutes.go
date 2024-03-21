package routes

import (
	"github.com/andrres017/technical-test/controllers"

	"github.com/labstack/echo/v4"
)

func ChallengeRoutes(e *echo.Group) {
	e.POST("/challenges", controllers.CreateChallengeHandler)
	e.GET("/challenges", controllers.FetchChallengesHandler)
	e.GET("/challenges/:id", controllers.GetChallengeByIDHandler)
	e.PUT("/challenges/:id", controllers.UpdateChallengeHandler)
	e.DELETE("/challenges/:id", controllers.DeleteChallengeHandler)
}
