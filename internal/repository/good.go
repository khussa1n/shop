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
	query := `
		SELECT  g.id, 
				g.name, 
				o.number, 
				go.good_count, 
				MAX(s.name) FILTER (WHERE gs.main_or_additional = 'главный') AS main_shelf,
				ARRAY_AGG(s.name) FILTER (WHERE gs.main_or_additional = 'дополнительный') AS additional_shelves
		FROM goods g
		JOIN goods_orders go ON g.id = go.good_id
		JOIN orders o ON go.order_id = o.id
		LEFT JOIN goods_shelves gs ON g.id = gs.good_id
		LEFT JOIN shelves s ON gs.shelf_id = s.id
		WHERE
	`

	for i := 0; i < len(orderNumbers); i++ {
		if i == 0 {
			query += " o.number = '" + orderNumbers[i] + "'"
		} else {
			query += " OR o.number = '" + orderNumbers[i] + "'"
		}
	}

	query += " GROUP BY g.id, g.name, o.number, go.good_count;"

	rows, err := g.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(dto.AllGoodsByOrders)
	for rows.Next() {
		var good dto.GoodWithOrders
		var mainShelf string
		if err = rows.Scan(&good.Good.ID, &good.Good.Name, &good.OrderNumber, &good.GoodsCount, &mainShelf, &good.AdditionalShelves); err != nil {
			return nil, err
		}

		result[mainShelf] = append(result[mainShelf], good)
	}

	return result, nil
}

func NewGood(pool *pgxpool.Pool) *GoodRepoImpl {
	return &GoodRepoImpl{pool: pool}
}
