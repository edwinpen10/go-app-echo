package routes

import (
	"log"
	"myBE/app/controllers"
	middleware "myBE/app/middlewares"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// uc := controllers.UserController{}

	e.GET("/user", controllers.GetUser)
	e.POST("/register", controllers.RegisterUser)
	e.POST("/login", controllers.Login)

	// isLogin := middleware.JWT([]byte(os.Getenv("JWT_SECRET")))
	// e.GET("/profile/:id", controllers.GetProfile, isLogin)

	r := e.Group("/profile")
	// config := echojwt.Config{
	// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	// 		return new(models.JwtCustomClaims)
	// 	},
	// 	SigningKey: []byte(os.Getenv("JWT_SECRET")),
	// }
	// r.Use(echojwt.WithConfig(config))
	r.Use(middleware.CustomJWTMiddleware())
	r.GET("/:id", controllers.GetProfile)
	r.PUT("/update/:id", controllers.UpdateUser)

	g := e.Group("/admin")
	g.GET("/", controllers.Index)
}
