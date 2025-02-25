package routes

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func LoadRoutes() *echo.Echo {
	router := echo.New()

	router.Use(echoprometheus.NewMiddleware("todo-rest-api"))
	router.GET("/metrics", echoprometheus.NewHandler())

	apiGroup := router.Group("/api")
	loadAuthRoutes(apiGroup)
	loadDocumentationRoutes(apiGroup)

	userGroup := apiGroup.Group("/user/:userId")
	loadTaskRoutes(userGroup)
	loadCollectionRoutes(userGroup)

	return router
}
