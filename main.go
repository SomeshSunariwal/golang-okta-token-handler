package main

import (
	"net/http"
	"os"

	"github.com/SomeshSunariwal/golang-okta-token-handler/handler"
	"github.com/SomeshSunariwal/golang-okta-token-handler/tokenMiddleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type LocalHandler struct {
	handler      handler.Handler
	tokenHandler tokenMiddleware.TokenHandlerMiddleware
}

func main() {

	Handler := LocalHandler{}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("/data", Handler.handler.GetData, Handler.tokenHandler.TokenHandler)

	// Testing CORS 
	HOST := os.Getenv("HOST")
	if HOST == "" {
		HOST="http://localhost:3000"
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{HOST},
		AllowMethods: []string{http.MethodGet},
		// AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":" + PORT))
}
