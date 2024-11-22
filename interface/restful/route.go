package restful

import (
	echo "github.com/labstack/echo/v4"
)

func NewRestfulServer() {
	e := echo.New()

	CoreExecutor(e.Group("/execute"))
	e.Start(":8080")
}
