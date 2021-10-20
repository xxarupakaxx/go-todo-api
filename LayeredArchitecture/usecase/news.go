package usecase
// validateとかpasswordのhash化とかをここでやる
import (
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain/repository"
)

type newsUseCase struct {
	newsRepository repository.NewsRepository
}

func (n newsUseCase) GetNews(id int) (*domain.News, error) {
	news,err := n.newsRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return news,nil
}

func (n newsUseCase) GetAllNews() ([]domain.News, error) {
	panic("implement me")
}

func (n newsUseCase) AddNews(news domain.News) error {
	panic("implement me")
}

func (n newsUseCase) GetAllNewsByFilter(status string) ([]domain.News, error) {
	panic("implement me")
}

func (n newsUseCase) RemoveNews(id int) error {
	panic("implement me")
}

func (n newsUseCase) UpdateNews(news *domain.News) error {
	panic("implement me")
}

func (n newsUseCase) GetNewsByTopic(slug string) ([]*domain.News, error) {
	panic("implement me")
}

type NewsUseCase interface {
	GetNews(id int) (*domain.News,error)
	GetAllNews() ([]domain.News,error)
	AddNews(news domain.News)error
	GetAllNewsByFilter(status string) ([]domain.News, error)
	RemoveNews(id int) error
	UpdateNews(news *domain.News) error
	GetNewsByTopic(slug string) ([]*domain.News, error)
}

func newNewsUseCase(newsRepository repository.NewsRepository) NewsUseCase {
	return &newsUseCase{newsRepository: newsRepository}
}

