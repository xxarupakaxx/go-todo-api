package mysql

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/xxarupakaxx/go-todo-api/go-clean-arch/article/repository"
	articleMysqlRepo "github.com/xxarupakaxx/go-todo-api/go-clean-arch/article/repository/mysql"
	"github.com/xxarupakaxx/go-todo-api/go-clean-arch/domain"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	db,mock,err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub databbase connection",err)
	}

	mockArticles := []domain.Article{
		{
			ID:        1,
			Title:     "Title 1",
			Content:   "Content 1",
			Author:    domain.Author{ID: 1},
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "Title 2",
			Content:   "Content 2",
			Author:    domain.Author{ID: 2},
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id","title","content","author_id","updated_at","created_at"}).
		AddRow(mockArticles[0].ID,mockArticles[0].Title,mockArticles[0].Content,mockArticles[0].Author.ID,mockArticles[0].UpdatedAt,mockArticles[0].Content).
		AddRow(mockArticles[1].ID,mockArticles[1].Title,mockArticles[1].Content,mockArticles[1].Author.ID,mockArticles[1].UpdatedAt,mockArticles[1].Content)

	query := "SELECT id,title,content,author_id,updated_at, created_at FROM article WHERE created_at > \\? ORDER BY created_at LIMIT \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := articleMysqlRepo.NewMysqlArticleRepository(db)
	cursor := repository.EncodeCursor(mockArticles[1].CreatedAt)
	num :=int64(2)
	list,nextCursor,err := a.Fetch(context.TODO(),cursor,num)
	assert.NotEmpty(t, nextCursor)
	assert.NotEmpty(t, err)
	assert.Len(t, list,2)
}