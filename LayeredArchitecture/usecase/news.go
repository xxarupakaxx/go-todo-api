package usecase
// validateとかpasswordのhash化とかをここでやる
import (
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain/repository"
)

type newsUseCase struct {
	newsRepository repository.NewsRepository
}

func (n *newsUseCase) Get(id int) (*domain.News, error) {
	news,err := n.newsRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return news,nil
}

func (n *newsUseCase) GetAll() ([]domain.News, error) {
	news,err := n.newsRepository.GetAll()
	if err!=nil {
		return nil, err
	}
	return news,nil
}

func (n *newsUseCase) GetBySlug(slug string) ([]*domain.News, error) {
	news,err := n.newsRepository.GetBySlug(slug)
	if err != nil {
		return nil, err
	}
	return news,nil
}

func (n *newsUseCase) GetAllByStatus(status string) ([]domain.News, error) {
	news,err := n.newsRepository.GetAllByStatus(status)
	if err != nil {
		return nil, err
	}
	return news,err
}

func (n *newsUseCase) Save(news *domain.News) error {
	err := n.newsRepository.Save(news)
	if err != nil {
		return err
	}
	return nil
}

func (n *newsUseCase) Remove(id int) error {
	err := n.newsRepository.Remove(id)
	if err != nil {
		return err
	}
	return nil
}

func (n *newsUseCase) Update(news *domain.News) error {
	err := n.newsRepository.Update(news)
	if err != nil {
		return err
	}
	return nil
}

type NewsUseCase interface {
	Get(id int) (*domain.News,error)
	GetAll() ([]domain.News,error)
	GetBySlug(slug string)([]*domain.News,error)
	GetAllByStatus(status string)([]domain.News,error)
	Save(news *domain.News) error
	Remove(id int) error
	Update(news *domain.News) error
}

func newNewsUseCase(newsRepository repository.NewsRepository) NewsUseCase {
	return &newsUseCase{newsRepository: newsRepository}
}

