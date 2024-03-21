package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/khussa1n/shop/internal/dto"
)

type GoodRepoImpl struct {
	pool *pgxpool.Pool
}

func (g *GoodRepoImpl) GetAllByOrders(orderNumbers ...string) (dto.AllGoodsByOrders, error) {
	query1 := `
		SELECT g.id, g.name, o.number, go.good_count, s.name
		FROM goods g
		JOIN goods_orders go ON g.id = go.good_id
		JOIN orders o ON go.order_id = o.id
		LEFT JOIN goods_shelves gs ON g.id = gs.good_id
		LEFT JOIN shelves s ON gs.shelf_id = s.id
		WHERE gs.main_or_additional = 'главный' AND (
	`

	for i := 0; i < len(orderNumbers); i++ {
		if i == 0 {
			query1 += "o.number = '" + orderNumbers[i] + "'"
		} else {
			query1 += " OR o.number = '" + orderNumbers[i] + "'"
		}
	}

	query1 += ") GROUP BY g.id, g.name, o.number, go.good_count, s.name;"

	rows1, err := g.pool.Query(context.Background(), query1)
	if err != nil {
		return nil, err
	}
	defer rows1.Close()

	allGoods := make(dto.AllGoodsByOrders)

	for rows1.Next() {
		var good dto.GoodWithOrders
		var shelfName string
		if err = rows1.Scan(&good.Good.ID, &good.Good.Name, &good.OrderNumber, &good.GoodsCount, &shelfName); err != nil {
			return nil, err
		}

		good.AdditionalShelves = make([]string, 0)

		allGoods[shelfName] = append(allGoods[shelfName], good)
	}

	query2 := `
		select s.name, gs.good_id from shelves s
		join goods_shelves gs on s.id = gs.shelf_id
		where gs.main_or_additional = 'дополнительный';
	`

	rows2, err := g.pool.Query(context.Background(), query2)
	if err != nil {
		return nil, err
	}
	defer rows2.Close()

	for rows2.Next() {
		var goodID int64
		var additionalShelf string
		if err = rows2.Scan(&additionalShelf, &goodID); err != nil {
			return nil, err
		}

		for shelfName, goods := range allGoods {
			for idx, good := range goods {
				if good.Good.ID == goodID {
					allGoods[shelfName][idx].AdditionalShelves = append(allGoods[shelfName][idx].AdditionalShelves, additionalShelf)
					break
				}
			}
		}
	}

	return allGoods, nil
}

func NewGood(pool *pgxpool.Pool) *GoodRepoImpl {
	return &GoodRepoImpl{pool: pool}
}
