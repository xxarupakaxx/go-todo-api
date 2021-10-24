package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/config"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain/repository"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/infrastructure/persistence"
	interfaces "github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/interfaces/handler"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/usecase"
	"log"
)

func InjectDB() *gorm.DB {
	conn,err := config.ConnectDB()
	if err != nil {
		log.Fatalln(err)
	}else {
		fmt.Println("Connect success")
	}
	return conn
}

func InjectNewsRepository() repository.NewsRepository {
	conn := InjectDB()
	return persistence.NewNewsRepositoryWithRDB(conn)
}

func InjectTopicRepository() repository.TopicRepository {
	conn := InjectDB()
	return persistence.NewTopicRepositoryWithRDB(conn)
}

func InjectTopicUseCas() usecase.TopicUseCase {
	topicRepo := InjectTopicRepository()
	return usecase.NewTopicUseCase(topicRepo)
}
func InjectNewsUseCase() usecase.NewsUseCase {
	newsRepo := InjectNewsRepository()
	return usecase.NewNewsUseCase(newsRepo)
}

func InjectNewsHandler() *interfaces.NewsHandler {
	newsUseCase := InjectNewsUseCase()
	return interfaces.NewNewsHandler(newsUseCase)
}

func InjectTopicHandler() *interfaces.TopicHandler {
	topicUseCase := InjectTopicUseCas()
	return interfaces.NewTopicHandler(topicUseCase)
}