package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/khussa1n/shop/internal/model"
	"strconv"
)

type GoodRepoImpl struct {
	pool *pgxpool.Pool
}

func (g *GoodRepoImpl) GetAllGoodsByIds(ids ...int64) ([]model.Goods, error) {
	var goods []model.Goods
	query := `
		select * from goods g
		where`

	for i := 0; i < len(ids); i++ {
		if i == 0 {
			query += " g.id = " + strconv.FormatInt(ids[i], 10)
		} else {
			query += " OR g.id = " + strconv.FormatInt(ids[i], 10)
		}
	}

	err := pgxscan.Select(context.Background(), g.pool, &goods, query)
	if err != nil {
		return nil, err
	}

	return goods, nil
}

func (g *GoodRepoImpl) GetOrdersByNumbers(numbers ...string) ([]model.OrdersByGoods, error) {
	var result []model.OrdersByGoods
	query := `
		select o.number, go.good_id, go.good_count from orders o
		join goods_orders go on o.id = go.order_id
		where`

	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			query += " o.number = " + numbers[i]
		} else {
			query += " OR o.number = " + numbers[i]
		}
	}

	err := pgxscan.Select(context.Background(), g.pool, &result, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (g *GoodRepoImpl) GetShelvesByGoods(ids ...int64) ([]model.ShelvesByGoods, error) {
	var result []model.ShelvesByGoods
	query := `
		select gs.good_id,
			   MAX(s.name) FILTER (WHERE gs.main_or_additional = 'главный') AS main_shelf,
			   ARRAY_AGG(s.name) FILTER (WHERE gs.main_or_additional = 'дополнительный') AS additional_shelves
		FROM goods_shelves gs
		join shelves s on gs.shelf_id = s.id
		WHERE`

	for i := 0; i < len(ids); i++ {
		if i == 0 {
			query += " gs.good_id = " + strconv.FormatInt(ids[i], 10)
		} else {
			query += " OR gs.good_id = " + strconv.FormatInt(ids[i], 10)
		}
	}

	query += " GROUP BY gs.good_id;"

	err := pgxscan.Select(context.Background(), g.pool, &result, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewGood(pool *pgxpool.Pool) *GoodRepoImpl {
	return &GoodRepoImpl{pool: pool}
}
