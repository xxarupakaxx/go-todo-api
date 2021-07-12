package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/xxarupakaxx/go-todo-api/src/controller"
)

func main() {

	e:=echo.New()
	v1:=e.Group("/todo/api/v1")
	{
		v1.GET("/tasks",controller.TaskGET)
		v1.POST("/tasks",controller.TaskPOST)
		v1.PATCH("/tasks/:id",controller.TaskPATCH)
		v1.DELETE("/tasks/:id",controller.TaskDELETE)
	}

	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":9000"))
}
