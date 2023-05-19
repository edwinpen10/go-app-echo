package routes

import (
	"myBE/app/controllers"
	models "myBE/app/models/users"
	"os"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo) {
	// uc := controllers.UserController{}

	e.GET("/user", controllers.GetUser)
	e.POST("/register", controllers.RegisterUser)
	e.POST("/login", controllers.Login)

	// isLogin := middleware.JWT([]byte(os.Getenv("JWT_SECRET")))
	// e.GET("/profile/:id", controllers.GetProfile, isLogin)

	r := e.Group("/profile")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("/:id", controllers.GetProfile)
	r.PUT("/update/:id", controllers.UpdateUser)

	g := e.Group("/admin")
	g.GET("/", controllers.Index)
}
