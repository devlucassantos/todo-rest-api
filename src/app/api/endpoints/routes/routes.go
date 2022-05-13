package routes

import "github.com/labstack/echo/v4"

func LoadRoutes() *echo.Echo {
	router := echo.New()

	apiGroup := router.Group("/api")
	loadAuthRoutes(apiGroup)

	userGroup := apiGroup.Group("/user/:userId")
	loadTaskRoutes(userGroup)
	loadCollectionRoutes(userGroup)

	return router
}
