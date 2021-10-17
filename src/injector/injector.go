package injector

import (
	"github.com/xxarupakaxx/go-todo-api/src/domain/repository"
	"github.com/xxarupakaxx/go-todo-api/src/handler"
	"github.com/xxarupakaxx/go-todo-api/src/infra"
	"github.com/xxarupakaxx/go-todo-api/src/usecase"
)

func InjectDB() infra.SqlHandler {
	sqlhandler := infra.NewSqlHandler()
	return *sqlhandler
}

func InjectTodoRepository() repository.TodoRepository {
	sqlHandler := InjectDB()
	return infra.NewTodoRepository(sqlHandler)
}

func InjectTodoUsecase() usecase.TodoUsecase {
	todoRepo := InjectTodoRepository()
	return usecase.NewTodoUsecase(todoRepo)
}

func InjectTodoHandler() handler.TodoHandler {
	return handler.NewTodoHandler(InjectTodoUsecase())
}
