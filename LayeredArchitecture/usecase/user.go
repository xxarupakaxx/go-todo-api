package usecase
// validateとかpasswordのhash化とかをここでやる
import (
	"database/sql"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain"
	"github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/domain/repository"
)

type userUseCase struct {
	userRepository repository.UserRepository
}

func (uu *userUseCase) GetByUserID(DB *sql.DB, userID string) (*domain.User, error) {
	user,err := uu.userRepository.GetByUserID(DB,userID)
	if err != nil {
		return nil, err
	}
	return user,nil
}

func (uu *userUseCase) Insert(DB *sql.DB,  name, email string) error {
	err := uu.userRepository.Insert(DB,name,email)
	if err != nil {
		return err
	}
	return nil
}

func NewUserUserCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{userRepository: ur}
}

type UserUseCase interface {
	GetByUserID(DB *sql.DB,userID string)(*domain.User,error)
	Insert(DB *sql.DB,name,email string) error
}

