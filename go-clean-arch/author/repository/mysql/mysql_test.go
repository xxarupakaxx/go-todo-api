package mysql_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	repository "github.com/xxarupakaxx/go-todo-api/go-clean-arch/author/repository/mysql"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow(1, "traP", time.Now(), time.Now())

	query := "SELECT id,name,created_at,updated_at FROM author WHERE id =\\?"

	prep := mock.ExpectPrepare(query)
	userID := int64(1)
	prep.ExpectQuery().WithArgs(1).WillReturnRows(rows)

	a := repository.NewMysqlAuthorRepository(db)

	anArticle, err := a.GetByID(context.TODO(), userID)

	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}
