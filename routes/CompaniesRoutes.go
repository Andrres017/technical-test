package routes

import (
	"github.com/andrres017/technical-test/controllers"

	"github.com/labstack/echo/v4"
)

func CompaniesRoutes(e *echo.Group) {
	e.POST("/companies", controllers.CreateCompanyHandler)
	e.GET("/companies", controllers.FetchCompaniesHandler)
	e.GET("/companies/:id", controllers.GetCompanyByIDHandler)
	e.PUT("/companies/:id", controllers.UpdateCompanyHandler)
	e.DELETE("/companies/:id", controllers.DeleteCompanyHandler)
}
