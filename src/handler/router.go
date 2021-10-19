package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouting(e *echo.Echo, todoHandler TodoHandler) {

	e.Use(middleware.Logger())

	e.GET("/", todoHandler.View())
	e.GET("/search", todoHandler.Search())
	e.POST("/create", todoHandler.Add())
	e.POST("/edit", todoHandler.Edit())
}
