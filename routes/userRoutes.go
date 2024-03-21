package routes

import (
	"github.com/andrres017/technical-test/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	e.POST("/users", controllers.CreateUser)
	e.GET("/users", controllers.GetUsersPaginated)
	e.GET("/users/:id", controllers.GetUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
}
