package controllers

import (
	"fmt"
	"net/http"

	"github.com/andrres017/technical-test/usecases"
	"github.com/andrres017/technical-test/utils"
	"github.com/labstack/echo/v4"
)

func HandleGPTAutoFill(c echo.Context) error {
	gptUseCase := usecases.NewGPTAutoFillUseCase()
	tables := []string{"Challenge", "Companies", "Program", "User"}

	for _, model := range tables {
		err := gptUseCase.GenerateAndSaveFakeData(model)
		if err != nil {
			// Puedes decidir cómo manejar el error. Aquí simplemente lo imprimimos.
			fmt.Printf("Error processing model %s: %v\n", model, err)
		}
	}

	fmt.Println("All data generation and saving complete.")
	return utils.ApiResponse(c, http.StatusOK, "Success", "Fake data generated", nil)
}
