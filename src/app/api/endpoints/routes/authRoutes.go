package routes

import (
	"github.com/labstack/echo/v4"
	"todo/src/app/api/endpoints/handlers"
)

func loadAuthRoutes(group *echo.Group) {
	authGroup := group.Group("/auth")

	authHandler := handlers.NewAuthHandler()

	authGroup.POST("/signup", authHandler.SignUp)
	authGroup.POST("/signin", authHandler.SignIn)
}
