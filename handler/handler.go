package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (handler Handler) GetData(context echo.Context) error {

	data := []map[string]string{
		{"Name": "Value1",},
		{"Name": "Value2",},
		{"Name": "Value3",},
		{"Name": "Value4",},
		{"Name": "Value5",},
		{"Name": "Value6",},
		{"Name": "Value7",},
	}

	return context.JSON(http.StatusOK, data)
}
