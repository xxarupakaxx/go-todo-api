package persistence

import (
	"database/sql"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain/repository"
)

type userPersistence struct {}

func (up *userPersistence) Insert(DB *sql.DB, name, email string) error {
	stmt,err := DB.Prepare("INSERT INTO users(account_name,address) VALUES(,?,?)")
	if err != nil {
		return err
	}
	_,err = stmt.Exec(name,email)
	return err
}

func (up *userPersistence) GetByUserID(DB *sql.DB, userID string) (*domain.User, error) {
	row := DB.QueryRow("SELECT * FROM users WHERE id = ?",userID)
	return convertToUser(row)
}

func convertToUser(row *sql.Row) (*domain.User, error) {
	user := domain.User{}
	err := row.Scan(&user.UserID,&user.Name,&user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil,nil
		}
		return nil, err
	}
	return &user,nil
}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}
