package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (handler Handler) GetData(context echo.Context) error {

	data := map[string]string{
		"Name1": "Value1",
		"Name2": "Value2",
		"Name3": "Value3",
		"Name4": "Value4",
		"Name5": "Value5",
		"Name6": "Value6",
		"Name7": "Value7",
	}

	return context.JSON(http.StatusOK, data)
}
