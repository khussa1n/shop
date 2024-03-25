package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/khussa1n/shop/internal/model"
	"strings"
)

type OrderRepoImpl struct {
	pool *pgxpool.Pool
}

func (o *OrderRepoImpl) GetAllByIds(numbers ...string) ([]model.Orders, error) {
	var orders []model.Orders
	query := `
		select * from orders
		where number in (`

	placeholders := make([]string, len(numbers))
	for i, id := range numbers {
		placeholders[i] = id
	}

	query += strings.Join(placeholders, ", ")
	query += ")"

	err := pgxscan.Select(context.Background(), o.pool, &orders, query)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func NewOrder(pool *pgxpool.Pool) *OrderRepoImpl {
	return &OrderRepoImpl{pool: pool}
}
