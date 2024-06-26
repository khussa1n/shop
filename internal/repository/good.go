package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/khussa1n/shop/internal/model"
	"strconv"
	"strings"
)

type GoodRepoImpl struct {
	pool *pgxpool.Pool
}

func (g *GoodRepoImpl) GetAllByIds(ids ...int64) ([]model.Goods, error) {
	var goods []model.Goods
	query := `
		select * from goods g
		where g.id in (`

	placeholders := make([]string, len(ids))
	for i, id := range ids {
		placeholders[i] = strconv.FormatInt(id, 10)
	}

	query += strings.Join(placeholders, ", ")
	query += ")"

	err := pgxscan.Select(context.Background(), g.pool, &goods, query)
	if err != nil {
		return nil, err
	}

	return goods, nil
}

func NewGood(pool *pgxpool.Pool) *GoodRepoImpl {
	return &GoodRepoImpl{pool: pool}
}
