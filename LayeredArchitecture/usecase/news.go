package usecase

// validateとかpasswordのhash化とかをここでやる
import (
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain/repository"
)

type newsUseCase struct {
	newsRepository repository.NewsRepository
}

func (n *newsUseCase) GetNews(id int) (*domain.News, error) {
	news, err := n.newsRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (n *newsUseCase) GetAllNews() ([]*domain.News, error) {
	news, err := n.newsRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (n *newsUseCase) AddNews(news *domain.News) error {
	if err := n.newsRepository.Save(news); err != nil {
		return err
	}
	return nil
}

func (n *newsUseCase) GetAllNewsByFilter(status string) ([]*domain.News, error) {
	news, err := n.newsRepository.GetAllByStatus(status)
	if err != nil {
		return nil, err
	}
	return news, err
}

func (n *newsUseCase) RemoveNews(id int) error {
	if err := n.RemoveNews(id); err != nil {
		return err
	}
	return nil
}

func (n *newsUseCase) UpdateNews(news *domain.News, id int) error {
	if err := n.newsRepository.Update(news); err != nil {
		return err
	}
	news.ID = (uint)(id)
	return nil
}

func (n *newsUseCase) GetNewsByTopic(slug string) ([]*domain.News, error) {
	news, err := n.newsRepository.GetBySlug(slug)
	if err != nil {
		return nil, err
	}
	return news, err
}

type NewsUseCase interface {
	GetNews(id int) (*domain.News, error)
	GetAllNews() ([]*domain.News, error)
	AddNews(news *domain.News) error
	GetAllNewsByFilter(status string) ([]*domain.News, error)
	RemoveNews(id int) error
	UpdateNews(news *domain.News, id int) error
	GetNewsByTopic(slug string) ([]*domain.News, error)
}

func NewNewsUseCase(newsRepository repository.NewsRepository) NewsUseCase {
	return &newsUseCase{newsRepository: newsRepository}
}
