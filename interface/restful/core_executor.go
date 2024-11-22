package restful

import (
	"edmk/application/core"
	"edmk/application/core/dto"
	"edmk/application/interfaces"

	"github.com/labstack/echo/v4"
)

func CoreExecutor(eg *echo.Group) {
	var coreService interfaces.CoreService = core.NewCoreService()

	eg.POST("/:kernel_id/:command", func(c echo.Context) error {
		var requestBody dto.BillerRequest
		if err := c.Bind(&requestBody); err != nil {
			return c.JSON(400, map[string]interface{}{
				"error": "invalid request body",
			})
		}

		kernelID := c.Param("kernel_id")
		command := c.Param("command")

		responseBody, err := coreService.BillerExecute(c.Request().Context(), kernelID, command, requestBody)
		if err != nil {
			return c.JSON(500, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.JSON(200, responseBody)
	})
}
