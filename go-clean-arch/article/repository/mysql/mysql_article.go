package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xxarupakaxx/go-todo-api/go-clean-arch/article/repository"
	"github.com/xxarupakaxx/go-todo-api/go-clean-arch/domain"
)

type mysqlArticleRepository struct {
	Conn *sql.DB
}

func (m *mysqlArticleRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Article, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.Article, 0)
	for rows.Next() {
		t := domain.Article{}
		authorID := int64(0)
		err = rows.Scan(&t.ID, &t.Title, &t.Content, &authorID, &t.UpdatedAt, &t.CreatedAt)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		t.Author = domain.Author{ID: authorID}
		result = append(result, t)

	}
	return result, err
}
func (m *mysqlArticleRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Article, nextCursor string, err error) {
	query := `SELECT id,title,content,author_id,updated_at,created_at FROM article WHERE created_at = ? ORDER BY created_at LIMIT ?`
	decodedCursor, err := repository.DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", err

	}

	res, err = m.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}
	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(res[len(res)-1].CreatedAt)
	}
	return
}

func (m *mysqlArticleRepository) GetByID(ctx context.Context, id int64) (res domain.Article, err error) {
	query := `SELECT id,title,content,author_id,updated_at,created_at FROM article WHERE ID =?`
	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Article{}, err
	}
	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}

func (m *mysqlArticleRepository) GetByTitle(ctx context.Context, title string) (res domain.Article, err error) {
	query := `SELECT id,title,content,author_id, updated_at,created_at FROM article WHERE title = ?`
	list, err := m.fetch(ctx, query, title)
	if err != nil {
		return
	}
	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}

func (m *mysqlArticleRepository) Update(ctx context.Context, ar *domain.Article) (err error) {
	query := `UPDATE article SET title =?,content = ?,author_id=?,updated_at=? WHERE ID =?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, ar.Title, ar.Content, ar.Author.ID, ar.UpdatedAt, ar.ID)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("weird Behavior. Total affected : %d", affect)
		return
	}
	return
}

func (m *mysqlArticleRepository) Store(ctx context.Context, a *domain.Article) (err error) {
	query := `INSERT article SET title = ?, content=?,author_id =?,updated_at=?,created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, a.Title, a.Content, a.Author.ID, a.UpdatedAt, a.CreatedAt)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	a.ID = lastID
	return

}

func (m *mysqlArticleRepository) Delete(ctx context.Context, id int64) (err error) {
	query := `DELETE FROM article WHERE id = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected != 1 {
		err = fmt.Errorf("weird Behavior. Total Affected: %d", rowsAffected)
		return
	}
	return
}

func NewMysqlArticleRepository(conn *sql.DB) domain.ArticleRepository {
	return &mysqlArticleRepository{Conn: conn}
}
