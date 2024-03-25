package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/khussa1n/shop/internal/model"
	"strconv"
	"strings"
)

type GoodShelfRepoImpl struct {
	pool *pgxpool.Pool
}

func (o *GoodShelfRepoImpl) GetAllByGoodIds(ids ...int64) ([]model.GoodsShelves, error) {
	var result []model.GoodsShelves
	query := `
		select * from goods_shelves
		where good_id in (`

	placeholders := make([]string, len(ids))
	for i, id := range ids {
		placeholders[i] = strconv.FormatInt(id, 10)
	}

	query += strings.Join(placeholders, ", ")
	query += ")"

	err := pgxscan.Select(context.Background(), o.pool, &result, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewGoodShelf(pool *pgxpool.Pool) *GoodShelfRepoImpl {
	return &GoodShelfRepoImpl{pool: pool}
}
