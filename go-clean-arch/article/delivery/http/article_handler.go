package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/xxarupakaxx/go-todo-api/go-clean-arch/domain"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ArticleHandler struct {
	AUsecase domain.ArticleUsecase
}

func NewArticleHandler(e *echo.Echo, us domain.ArticleUsecase) {
	handler := &ArticleHandler{AUsecase: us}
	e.GET("/articles", handler.FetchArticle)
	e.POST("/articles",handler.Store)
	e.GET("/articles/:id",handler.GetByID)
	e.DELETE("/articles/:id",handler.Delete)
}

func (a *ArticleHandler) FetchArticle(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()

	listAr, nextCursor, err := a.AUsecase.Fetch(ctx, cursor, int64(num))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}

func (a *ArticleHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	id := int64(idP)
	ctx := c.Request().Context()

	art, err := a.AUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)

}

func isRequestValid(m *domain.Article) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, err
}

func (a *ArticleHandler) Store(c echo.Context) error {
	var article domain.Article
	err := c.Bind(&article)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity,err.Error())
	}

	var ok bool
	if ok,err = isRequestValid(&article);!ok {
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	ctx := c.Request().Context()
	err = a.AUsecase.Store(ctx,&article)
	if err != nil {
		return c.JSON(getStatusCode(err),ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated,article)
}

func (a *ArticleHandler) Delete(c echo.Context) error {
	idP,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound,domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	err = a.AUsecase.Delete(ctx,id)
	if err != nil {
		return c.JSON(getStatusCode(err),ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case domain.ErrInternalServer:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
