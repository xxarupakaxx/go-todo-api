package persistence

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain/repository"
)

type NewsRepositoryImpl struct {
	Conn *gorm.DB
}

func (r *NewsRepositoryImpl) Get(id int) (*domain.News, error) {
	news := &domain.News{}
	if err := r.Conn.Preload("Topic").First(&news,id).Error;err !=nil{
		return nil, err
	}
	return news,nil
}

func (r *NewsRepositoryImpl) GetAll() ([]domain.News, error) {
	news :=[]domain.News{}
	if err := r.Conn.Preload("Topic").Find(&news).Error; err != nil {
		return nil, err
	}
	return news,nil
}

func (r *NewsRepositoryImpl) GetBySlug(slug string) ([]*domain.News, error) {
	rows, err := r.Conn.Raw("SELECT news.id, news.title, news.slug, news.content, news.status FROM `news_topics` LEFT JOIN news ON news_topics.news_id=news.id WHERE news_topics.topic_id=(SELECT id as topic_id FROM `topics` WHERE slug = ?)", slug).Rows()
	defer rows.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil,nil
		}
		return nil, err
	}

	us := make([]*domain.News,0)

	for rows.Next() {
		u := &domain.News{}
		err = rows.Scan(&u.ID, &u.Title, &u.Slug, &u.Content, &u.Status)

		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us,nil
}

func (r *NewsRepositoryImpl) GetAllByStatus(status string) ([]domain.News, error) {
	if status == "deleted" {
		news := []domain.News{}
		if err := r.Conn.Unscoped().Where("status = ?", status).Preload("Topic").Find(&news).Error; err != nil {
			return nil, err
		}

		return news, nil
	}

	news := []domain.News{}
	if err := r.Conn.Where("status = ?", status).Preload("Topic").Find(&news).Error; err != nil {
		return nil, err
	}

	return news, nil
}

func (r *NewsRepositoryImpl) Save(news *domain.News) error {
	if err := r.Conn.Save(&news).Error; err != nil {
		return err
	}
	return nil
}

func (r *NewsRepositoryImpl) Remove(id int) error {
	tx := r.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	news := domain.News{}
	if err := tx.First(&news, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	news.Status = "deleted"
	if err := tx.Save(&news).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&news).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *NewsRepositoryImpl) Update(news *domain.News) error {
	if err:=r.Conn.Model(&news).UpdateColumns(domain.News{
		Title:   news.Title,
		Slug:    news.Slug,
		Content: news.Content,
		Status:  news.Status,
		Topic:   news.Topic,
	}).Error;err !=nil{
		return err
	}
	return nil
}

func NewNewsRepositoryWithRDB(conn *gorm.DB) repository.NewsRepository {
	return &NewsRepositoryImpl{conn}
}