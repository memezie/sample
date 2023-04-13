package custommiddleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// server.goの可読性を上げるため、LoggerWithConfigをラップ
func LoggerWithConfig() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			value, _ := c.Get("customValueFromContext").(int)
			fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status, value)
			return nil
		},
	})
	// return middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	// })
}
