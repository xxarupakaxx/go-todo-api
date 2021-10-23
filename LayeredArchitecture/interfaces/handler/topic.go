package interfaces

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/usecase"
	"net/http"
	"strconv"
)

type TopicHandler struct {
	topicUseCase usecase.TopicUseCase
}

func newTopicHandler(topicUseCase usecase.TopicUseCase) *TopicHandler {
	return &TopicHandler{topicUseCase: topicUseCase}
}

func (th *TopicHandler) GetAllTopic() echo.HandlerFunc {
	return func(c echo.Context) error {
		topics,err :=th.topicUseCase.GetAllTopic()
		if err != nil {
			return fmt.Errorf("failed in Getting Alltopic:%w",err)
		}
		return c.JSON(http.StatusOK,topics)
	}
}

func (th *TopicHandler) GetTopic() echo.HandlerFunc {
	return func(c echo.Context) error {
		topicID,err :=strconv.Atoi(c.Param("id"))
		if err != nil {
			return fmt.Errorf("failed in chaging int :%w",err)
		}
		topic,err:=th.topicUseCase.GetTopic(topicID)
		if err != nil {
			return fmt.Errorf("failed in getting Topics:%w",err)
		}
		return c.JSON(http.StatusOK,topic)
	}
}

func (th *TopicHandler) CreateTopic() echo.HandlerFunc {
	return func(c echo.Context) error {
		type payload struct {
			Name string `json:"name"`
			Slug string `json:"slug"`
		}
		var p payload
		if err:=c.Bind(&p);err!=nil{
			return fmt.Errorf("failed in Bind :%w",err)
		}
		err := th.topicUseCase.AddTopic(p.Name,p.Slug)
		if err != nil {
			return fmt.Errorf("failed in AddTopic :%w",err)
		}
		return c.JSON(http.StatusCreated,nil)
	}
}

func (th *TopicHandler) UpdateTopic() echo.HandlerFunc {
	return func(c echo.Context) error {
		var topic domain.Topic
		if err := c.Bind(&topic); err != nil {
			return c.JSON(http.StatusNotFound,err)
		}
		topicID,err:=strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusNotFound,err)
		}
		err = th.topicUseCase.UpdateTopic(&topic,topicID)
		if err != nil {
			return c.JSON(http.StatusNotFound,err)
		}
		return c.JSON(http.StatusOK,nil)
	}
}

func (th *TopicHandler) RemoveTopic() echo.HandlerFunc {
	return func(c echo.Context) error {
		topicID,err := strconv.Atoi("id" )
		if err != nil {
			return c.JSON(http.StatusNotFound,err)
		}
		err = th.topicUseCase.Remove(topicID)
		if err != nil {
			return c.JSON(http.StatusNotFound,err)
		}
		return c.JSON(http.StatusOK,nil)
	}

}