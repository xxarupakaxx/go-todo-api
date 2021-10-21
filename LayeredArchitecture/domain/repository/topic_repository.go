package repository

import "github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"

type TopicRepository interface {
	Get(id int)(*domain.Topic,error)
	GetAll()([]*domain.Topic,error)
	Save(topic *domain.Topic) error
	Remove(id int) error
	Update(topic *domain.Topic) error
}