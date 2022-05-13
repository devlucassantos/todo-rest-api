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

	collectionHandler := handlers.NewCollectionHandler()
	taskHandler := handlers.NewTaskHandler()

	collectionGroup.POST("", collectionHandler.Create)
	collectionGroup.PUT("/:collectionId", collectionHandler.Update)
	collectionGroup.DELETE("/:collectionId", collectionHandler.Delete)
	collectionGroup.GET("", collectionHandler.FindAll)
	collectionGroup.GET("/:collectionId", collectionHandler.FindById)
	collectionGroup.GET("/:collectionId/tasks", taskHandler.FindByCollectionId)
}
