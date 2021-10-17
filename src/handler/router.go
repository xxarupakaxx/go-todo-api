package handler

import "github.com/labstack/echo"

func InitRouting(e *echo.Echo, todoHandler TodoHandler) {
	e.GET("/",todoHandler.View())
	e.GET("/search",todoHandler.Search())
	e.POST("/create",todoHandler.Add())
	e.POST("/edit",todoHandler.Edit())
}