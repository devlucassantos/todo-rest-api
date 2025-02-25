package routes

import (
	"github.com/labstack/echo/v4"
	"todo/src/app/api/endpoints/handlers"
	"todo/src/app/api/endpoints/middleware"
)

func loadTaskRoutes(group *echo.Group) {
	taskGroup := group.Group("/task")
	authMiddleware := middleware.NewAuthMiddleware()
	taskGroup.Use(authMiddleware.Authorize)

	taskHandler := handlers.NewTaskHandler()

	taskGroup.POST("", taskHandler.Create)
	taskGroup.PUT("/:taskId", taskHandler.Update)
	taskGroup.DELETE("/:taskId", taskHandler.Delete)
	taskGroup.GET("", taskHandler.FindAll)
}
