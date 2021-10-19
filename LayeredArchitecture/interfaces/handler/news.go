package interfaces

import (
	"github.com/labstack/echo"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/usecase"
	"net/http"
	"strconv"
)

type NewsHandler struct {
	newsUseCase usecase.NewsUseCase
}

func newNewsHandler(newsUseCase usecase.NewsUseCase) *NewsHandler {
	return &NewsHandler{newsUseCase: newsUseCase}
}

// Get GET /news/:id
func (nh *NewsHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.QueryParam("id")
		intID ,err := strconv.Atoi(id)
		if err != nil {
			return c.String(http.StatusInternalServerError,err.Error())
		}
		models, err := nh.newsUseCase.Get(intID)
		if err != nil {
			return c.String(http.StatusInternalServerError,err.Error())
		}
		return c.JSON(http.StatusOK,models)
	}
}
//GetAll GET /news
func (nh *NewsHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		news,err:= nh.newsUseCase.GetAll()
		if err != nil {
			return c.String(http.StatusBadRequest,err.Error())
		}
		return c.JSON(http.StatusOK,news)
	}
}

// NewsCreate POST /news
func (nh *NewsHandler) NewsCreate() echo.HandlerFunc {

}

// Update PUT /news/:id
func (nh *NewsHandler) Update() echo.HandlerFunc {

}

// Remove DELETE /news/:id
func (nh *NewsHandler) Remove() echo.HandlerFunc {

}