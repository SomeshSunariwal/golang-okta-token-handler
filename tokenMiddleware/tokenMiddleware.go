package tokenMiddleware

import (
	"github.com/labstack/echo/v4"
)

type TokenHandlerMiddleware struct {
}

func (handler TokenHandlerMiddleware) TokenHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(conext echo.Context) error {

		return next(conext)
	}
}
