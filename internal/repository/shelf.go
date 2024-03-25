package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/khussa1n/shop/internal/model"
	"strconv"
	"strings"
)

type ShelfRepoImpl struct {
	pool *pgxpool.Pool
}

func (s *ShelfRepoImpl) GetAllByIds(ids ...int64) ([]model.Shelves, error) {
	var shelves []model.Shelves
	query := `
		select * from shelves
		where id in (`

	placeholders := make([]string, len(ids))
	for i, id := range ids {
		placeholders[i] = strconv.FormatInt(id, 10)
	}

	query += strings.Join(placeholders, ", ")
	query += ")"

	err := pgxscan.Select(context.Background(), s.pool, &shelves, query)
	if err != nil {
		return nil, err
	}

	return shelves, nil
}

func NewShelf(pool *pgxpool.Pool) *ShelfRepoImpl {
	return &ShelfRepoImpl{pool: pool}
}
