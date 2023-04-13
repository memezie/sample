package agentController

import (
	"log"
	"net/http"
	"strconv"

	"example.com/mod/app/models"
	"github.com/labstack/echo/v4"
)

func GetAgents() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("in function")
		model := models.Agent{}
		agents, _ := model.GetAgents()
		return c.JSON(http.StatusOK, agents)
	}
}

func GetAgent() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		model := models.Agent{Id: id}
		model.GetAgent()
		return c.JSON(http.StatusOK, model)
	}
}
