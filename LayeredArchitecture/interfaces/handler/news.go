package interfaces

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"
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

// GetNews Get GET /news/:id
func (nh *NewsHandler) GetNews(c echo.Context) error {
		id := c.Param("id")
		intID ,err := strconv.Atoi(id)
		if err != nil {
			return c.String(http.StatusInternalServerError,err.Error())
		}
		models, err := nh.newsUseCase.GetNews(intID)
		if err != nil {
			return c.String(http.StatusInternalServerError,err.Error())
		}
		return c.JSON(http.StatusOK,models)
	}

// GetAllNews GetAll GET /news
func (nh *NewsHandler) GetAllNews(c echo.Context) error {
		status := c.QueryParam("status")
		if status =="draft" || status == "deleted" || status =="publish" {
			news,err := nh.newsUseCase.GetAllNewsByFilter(status)
			if err != nil {
				return c.JSON(http.StatusInternalServerError,err)
			}
			return c.JSON(http.StatusOK,news)
		}
		limit :=c.QueryParam("limit")
		page := c.QueryParam("page")

		if limit!=""&&page != "" {
			limit,_ := strconv.Atoi(limit)
			page,_:=strconv.Atoi(page)

			if limit != 0 && page != 0 {
				news,err :=nh.newsUseCase.GetAllNews()
				if err != nil {
					return fmt.Errorf("could not get All News %w",err)
				}
				return c.JSON(http.StatusOK,news)
			}

		}

		news,err := nh.newsUseCase.GetAllNews()
		if err != nil {
			return fmt.Errorf("could not get All News %w",err)
		}
		return c.JSON(http.StatusOK,news)
	}

// CreateNews NewsCreate POST /news
func (nh *NewsHandler) CreateNews(c echo.Context) error {

		var news domain.News
		if err:=c.Bind(&news);err!=nil{
			return fmt.Errorf("failed in bind News : %w",err)
		}
		err := nh.newsUseCase.AddNews(&news)
		return c.JSON(http.StatusOK,err)
	}

// UpdateNews Update PUT /news/:id
func (nh *NewsHandler) UpdateNews(c echo.Context) error{
		var news domain.News
		if err := c.Bind(&news); err != nil {
			return c.JSON(http.StatusNotFound,err)
		}
		newsID,err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusNotFound,err)
		}

		err = nh.newsUseCase.UpdateNews(&news,newsID)

		return c.JSON(http.StatusOK,err)
	}

// RemoveNews Remove DELETE /news/:id
func (nh *NewsHandler) RemoveNews(c echo.Context) error {
		newsID,err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return fmt.Errorf("failed in changing integer:%w",err)
		}
		err = nh.newsUseCase.RemoveNews(newsID)
		if err != nil {
			return fmt.Errorf("failed in Removing news :%w",err)
		}
		return c.JSON(http.StatusOK,nil)
	}