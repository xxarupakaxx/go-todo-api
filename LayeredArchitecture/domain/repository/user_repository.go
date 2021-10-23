package repository

import (
	"database/sql"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"
)

type UserRepository interface {
	Insert(DB *sql.DB, name, email string) error
	GetByUserID(DB *sql.DB, userID string) (*domain.User, error)
}
