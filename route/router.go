package route

import (
	"example.com/mod/app/controller/agentController"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {

	e := echo.New()

	api := e.Group("/api")
	{
		api.GET("/agents", agentController.GetAgents())
		api.GET("/agents/:id", agentController.GetAgent())
		// api.POST("/members", controller.PostMembers())
		// api.PUT("/members", controller.PutMembers())
		// api.DELETE("/members", controller.DeleteMembers())
	}
	return e
}
