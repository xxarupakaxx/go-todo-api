package usecase_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xxarupakaxx/go-todo-api/go-clean-arch/article/usecase"
	"github.com/xxarupakaxx/go-todo-api/go-clean-arch/domain"
	"github.com/xxarupakaxx/go-todo-api/go-clean-arch/domain/mock/domain"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	mockArticleRepo := new(mock_domain.MockArticleRepository)
	mockArticle := domain.Article{
		Title:   "Hello",
		Content: "Content",
	}

	mockListArticle := make([]domain.Article, 0)
	mockListArticle = append(mockListArticle, mockArticle)

	t.Run("success", func(t *testing.T) {
		mockArticleRepo.EXPECT().Fetch(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListArticle, "next-cursor", nil)
		mockAuthor := domain.Author{ID: 1, Name: "arupaka"}

		mockAuthorrepo := new(mock_domain.MockAuthorRepository)
		mockAuthorrepo.EXPECT().GetByID(mock.Anything, mock.AnythingOfType("int64")).Return(mockAuthor, nil)
		u := usecase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)
		num := int64(1)
		cursor := "12"
		list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)
		cursorExpected := "next-cursor"
		assert.Equal(t, cursorExpected, nextCursor)
		assert.NotEmpty(t, nextCursor)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListArticle))
	})

	t.Run("errror-failed", func(t *testing.T) {
		mockArticleRepo.EXPECT().Fetch(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(nil, "", errors.New("Unexpexted Error"))

		mockAuthorrepo := new(mock_domain.MockAuthorRepository)
		u := usecase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)
		num := int64(1)
		cursor := "12"
		list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)

		assert.Empty(t, nextCursor)
		assert.Error(t, err)
		assert.Len(t, list, 0)
	})
}
