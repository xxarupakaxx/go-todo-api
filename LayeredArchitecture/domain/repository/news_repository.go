package repository

import "github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"

type NewsRepository interface {
	Get(id int) (*domain.News, error)
	GetAll() ([]*domain.News, error)
	GetBySlug(slug string) ([]*domain.News, error)
	GetAllByStatus(status string) ([]*domain.News, error)
	Save(news *domain.News) error
	Remove(id int) error
	Update(news *domain.News) error
}
