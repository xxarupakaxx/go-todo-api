//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	interfaces "github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/interfaces/handler"
)

var (
	t = wire.Bind(new(interfaces.INewsHandler),new(*interfaces.NewsHandler))
)

func InjectionServer()  {
	wire.Build(
		interfaces.InitRouting())
}