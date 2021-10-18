package interfaces

import (
	"github.com/labstack/echo"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/usecase"
)

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func (uh *userHandler) HandleUserGet(c echo.Context) error {
	panic("implement me")
}

func (uh *userHandler) HandleUserSignup(c echo.Context) error {
	panic("implement me")
}

func newUserHandler(uu usecase.UserUseCase) *userHandler {
	return &userHandler{userUseCase: uu}
}

type UserHandler interface {
	HandleUserGet(c echo.Context) error
	HandleUserSignup(c echo.Context) error
}