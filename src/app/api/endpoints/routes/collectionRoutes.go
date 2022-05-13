package routes

import (
	"github.com/labstack/echo/v4"
	"todo/src/app/api/endpoints/handlers"
	"todo/src/app/api/endpoints/middleware"
)

func loadCollectionRoutes(group *echo.Group) {
	collectionGroup := group.Group("/collection")
	authMiddleware := middleware.NewAuthMiddleware()
	collectionGroup.Use(authMiddleware.Authorize)

	taskHandler := handlers.NewTaskHandler()

	collectionGroup.GET("/:collectionId/tasks", taskHandler.FindByCollectionId)
}
