package routes

import (
	"github.com/labstack/echo/v4"
	"todo/docs"
)

func loadDocumentationRoutes(group *echo.Group) {
	group.GET("/documentation/*", docs.WrapHandler)
}
