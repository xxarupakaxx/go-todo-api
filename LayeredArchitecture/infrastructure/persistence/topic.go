package persistence

import (
	"github.com/jinzhu/gorm"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain/repository"
)

type TopicRepositoryImpl struct {
	Conn *gorm.DB
}

func (t *TopicRepositoryImpl) Get(id int) (*domain.Topic, error) {
	topic := &domain.Topic{}
	if err:= t.Conn.Preload("News").First(&topic,id).Error;err!=nil{
		return nil, err
	}
	return topic,nil
}

func (t *TopicRepositoryImpl) GetAll() ([]*domain.Topic, error) {
	topics := []*domain.Topic{}
	if err:= t.Conn.Preload("News").Find(&topics).Error;err!=nil{
		return nil, err
	}

	return topics,nil
}

func (t *TopicRepositoryImpl) Save(topic *domain.Topic) error {
	if err := t.Conn.Save(&topic).Error;err!=nil{
		return err
	}
	return nil
}

func (t *TopicRepositoryImpl) Remove(id int) error {
	topic := &domain.Topic{}
	if err:= t.Conn.First(&topic,id).Error;err!=nil{
		return err
	}

	if err := t.Conn.Delete(&topic).Error; err != nil {
		return err
	}
	return nil
}

func (t *TopicRepositoryImpl) Update(topic *domain.Topic) error {
	if err := t.Conn.Model(&topic).UpdateColumns(domain.Topic{
		Name:  topic.Name,
		Slug:  topic.Slug,
	}).Error;err!=nil{
		return err
	}
	return nil
}

func NewTopicRepositoryWithRDB(conn *gorm.DB) repository.TopicRepository {
	return &TopicRepositoryImpl{Conn: conn}
}