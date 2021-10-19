package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/xxarupakaxx/go-todo-api/src/handler"
	"github.com/xxarupakaxx/go-todo-api/src/injector"
)

func main() {
	fmt.Println("server start ...")
	todoHandler := injector.InjectTodoHandler()
	e := echo.New()
	handler.InitRouting(e, todoHandler)
	e.Logger.Debug()
	e.Logger.Fatal(e.Start(":1449"))
}
