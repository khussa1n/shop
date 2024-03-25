package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/khussa1n/shop/internal/model"
	"strconv"
	"strings"
)

type GoodOrderRepoImpl struct {
	pool *pgxpool.Pool
}

func (o *GoodOrderRepoImpl) GetAllByOrderIds(ids ...int64) ([]model.GoodsOrders, error) {
	var result []model.GoodsOrders
	query := `
		select * from goods_orders
		where order_id in (`

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

func NewGoodOrder(pool *pgxpool.Pool) *GoodOrderRepoImpl {
	return &GoodOrderRepoImpl{pool: pool}
}
