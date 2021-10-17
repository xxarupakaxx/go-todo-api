package mock_repository

import (
	"github.com/golang/mock/gomock"
	"github.com/xxarupakaxx/go-todo-api/src/domain/model"
	"github.com/xxarupakaxx/go-todo-api/src/usecase"
	"reflect"
	"testing"
)

func TestView(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected []*model.Todo
	var err error

	mockSample := NewMockTodoRepository(ctrl)
	mockSample.EXPECT().FindAll().Return(expected, err)

	todoUsecase := usecase.NewTodoUsecase(mockSample)
	result, err := todoUsecase.View()

	if err != nil {
		t.Error("Actual FindAll() is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Error("Actual FindAll() is not same as expected")
	}

}

func TestSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected []*model.Todo
	var err error
	word := "test"

	mockSample := NewMockTodoRepository(ctrl)
	mockSample.EXPECT().Find(word).Return(expected, err)

	todoUsecase := usecase.NewTodoUsecase(mockSample)
	result, err := todoUsecase.Search(word)

	if err != nil {
		t.Error("Actual Find(word string) is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual Find(word string) is not same as expected")
	}
}

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected *model.Todo
	var err error

	mockSample := NewMockTodoRepository(ctrl)
	mockSample.EXPECT().Create(expected).Return(expected, err)

	todoUsecase := usecase.NewTodoUsecase(mockSample)
	err = todoUsecase.Add(expected)

	if err != nil {
		t.Error("Actual Find(word string) is not same as expected")
	}
}

func TestEdit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected *model.Todo
	var err error

	mockSample := NewMockTodoRepository(ctrl)
	mockSample.EXPECT().Update(expected).Return(expected, err)

	// mockを利用してtodoUsecase.Edit(todo *model.Todo)をテストする
	todoUsecase := usecase.NewTodoUsecase(mockSample)
	err = todoUsecase.Edit(expected)

	if err != nil {
		t.Error("Actual Find(word string) is not same as expected")
	}

}
