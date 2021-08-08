package tokenMiddleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type TokenHandlerMiddleware struct {
}

func (handler TokenHandlerMiddleware) TokenHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {

		// Getting Authorization header value
		AuthorizationHeader := context.Request().Header.Get("Authorization")
		if AuthorizationHeader == "" {
			return context.JSON(http.StatusUnauthorized, map[string]string{"message":"Unauthorized"})
		}
		tokenString := strings.Split(AuthorizationHeader, " ")[1]

		// Parsing Token without validation
		token, err := jwt.Parse(tokenString, nil)
		if token == nil {
			return context.JSON(http.StatusUnauthorized, map[string]string{"message":err.Error()})
		}
		claims, _ := token.Claims.(jwt.MapClaims)

		// Token Expire Check 
		utcTime := int64(claims["exp"].(float64))
		expireTime := time.Unix(utcTime, 0).Unix()
		if expireTime < 0 || expireTime < time.Now().Unix() {
			return context.JSON(http.StatusUnauthorized, map[string]string{"message":"Token Expire !"})
		}

		//only for testing. not corrrect from logical purpose
		if claims["name"] != "Somesh Sunariwal" || claims["sub"] != "someshsunariwal@gmail.com" || claims["iss"] != "https://dev-852842.okta.com/oauth2/default" {
			return context.JSON(http.StatusUnauthorized, map[string]string{"message":"Unauthorized"})
		}

		// Move to next middleware
		return next(context)
	}
}
