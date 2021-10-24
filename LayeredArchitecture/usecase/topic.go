package usecase

// validateとかpasswordのhash化とかをここでやる
import (
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain/repository"
)

type topicUseCase struct {
	topicRepository repository.TopicRepository
}

func NewTopicUseCase(topicRepository repository.TopicRepository) *topicUseCase {
	return &topicUseCase{topicRepository: topicRepository}
}

func (t *topicUseCase) GetTopic(id int) (*domain.Topic, error) {

	topic, err := t.topicRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return topic, err
}

func (t *topicUseCase) GetAllTopic() ([]*domain.Topic, error) {
	topics, err := t.topicRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return topics, err
}

func (t *topicUseCase) AddTopic(name, slug string) error {
	u := domain.Topic{
		Name: name,
		Slug: slug,
	}
	return t.topicRepository.Save(&u)
}

func (t *topicUseCase) UpdateTopic(topic *domain.Topic, id int) error {
	err := t.topicRepository.Update(topic)
	if err != nil {
		return err
	}
	topic.ID = uint(id)
	return nil
}

func (t *topicUseCase) Remove(id int) error {
	if err := t.topicRepository.Remove(id); err != nil {
		return err
	}
	return nil
}

func (t *topicUseCase) Update(topic *domain.Topic) error {
	if err := t.topicRepository.Update(topic); err != nil {
		return err
	}
	return nil
}

type TopicUseCase interface {
	GetTopic(id int) (*domain.Topic, error)
	GetAllTopic() ([]*domain.Topic, error)
	AddTopic(name, slug string) error
	UpdateTopic(topic *domain.Topic, id int) error
	Remove(id int) error
	Update(topic *domain.Topic) error
}
