package interfaces

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouting(topicHandler TopicHandler, newsHandler NewsHandler) {
	e := echo.New()
	e.Use(middleware.Logger())

	api := e.Group("/api")
	{
		apiNews := api.Group("/news")
		{
			apiNews.GET("", newsHandler.GetAllNews)
			apiNews.GET("/:id", newsHandler.GetNews)
			apiNews.POST("", newsHandler.CreateNews)
			apiNews.DELETE("/:id", newsHandler.RemoveNews)
			apiNews.PUT("/:id", newsHandler.UpdateNews)

		}

		apiTopic := api.Group("/topic")
		{
			apiTopic.GET("", topicHandler.GetAllTopic)
			apiTopic.GET("/:id", topicHandler.GetTopic)
			apiTopic.POST("", topicHandler.CreateTopic)
			apiTopic.DELETE("/:id", topicHandler.RemoveTopic)
			apiTopic.PUT("/:id", topicHandler.UpdateTopic)
		}
	}

}
