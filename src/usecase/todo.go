package usecase

import (
	"github.com/xxarupakaxx/go-todo-api/src/domain/model"
	"github.com/xxarupakaxx/go-todo-api/src/domain/repository"
)

type TodoUsecase interface {
	Search(string) (todo []*model.Todo,err error)
	View() (todo []*model.Todo,err error)
	Add(todo *model.Todo) (err error)
	Edit(todo *model.Todo)(err error)
}

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	todoUsercase := todoUsecase{todoRepo: todoRepo}
	return &todoUsercase
}

func (usecase *todoUsecase) Search(word string) (todo []*model.Todo, err error) {
	todo,err = usecase.todoRepo.Find(word)
	return
}

func (usecase *todoUsecase) View() (todo []*model.Todo, err error) {
	todo,err = usecase.todoRepo.FindAll()
	return 
}

func (usecase todoUsecase) Add(todo *model.Todo) (err error) {
	_,err = usecase.todoRepo.Create(todo)
	return 
}

func (usecase *todoUsecase) Edit(todo *model.Todo) (err error) {
	_, err = usecase.todoRepo.Update(todo)
	return
}

