package custommiddleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

type (
	BeforeAction struct {
		Text string `json:"text"`
	}
	AfterAction struct {
		Text string `json:"text"`
	}
)

func NewBeforeAction() *BeforeAction {
	return &BeforeAction{Text: "before action done."}
}

func (a *BeforeAction) Start(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println(a.Text)
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

func NewAfterAction() *BeforeAction {
	return &BeforeAction{Text: "after action done."}
}

func (a *AfterAction) Start(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
