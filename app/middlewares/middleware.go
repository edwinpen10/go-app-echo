package middleware

import (
	models "myBE/app/models/users"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ErrMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func CustomJWTMiddleware() echo.MiddlewareFunc {

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: func(c echo.Context, err error) error {
			message := err.Error()
			return echo.NewHTTPError(http.StatusUnauthorized, message).SetInternal(err)
		},
	}

	return echojwt.WithConfig(config)
}
