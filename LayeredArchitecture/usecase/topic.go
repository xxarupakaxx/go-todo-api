package usecase

import (
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain/repository"
)

type topicUseCase struct {
	topicRepository repository.TopicRepository
}

func (t *topicUseCase) Get(id int) (*domain.Topic, error) {
	topic,err := t.topicRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return topic,nil
}

func (t *topicUseCase) GetAll() ([]domain.Topic, error) {
	topics,err := t.topicRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return topics,nil
}

func (t *topicUseCase) Save(topic *domain.Topic) error {
	err := t.topicRepository.Save(topic)
	if err != nil {
		return err
	}
	return nil
}

func (t *topicUseCase) Remove(id int) error {
	err := t.topicRepository.Remove(id)
	if err != nil {
		return err
	}
	return err
}

func (t *topicUseCase) Update(topic *domain.Topic) error {
	err := t.topicRepository.Update(topic)
	if err != nil {
		return err
	}
	return nil
}

type TopicUseCase interface {
	Get(id int)(*domain.Topic,error)
	GetAll()([]domain.Topic,error)
	Save(topic *domain.Topic) error
	Remove(id int) error
	Update(topic *domain.Topic) error
}