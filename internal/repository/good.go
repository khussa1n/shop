package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/khussa1n/shop/internal/dto"
	"log"
)

type GoodRepoImpl struct {
	pool *pgxpool.Pool
}

func (g *GoodRepoImpl) GetAllByOrders(orderNumbers ...string) (dto.AllGoodsByOrders, error) {
	query := `
		SELECT s.name, g.id, g.name, o.number, go.good_count,
			   array_agg(distinct CASE WHEN gs.main_or_additional = 'дополнительный' THEN sh.name ELSE NULL END) AS additional_shelves
		FROM shelves s
		JOIN goods_shelves gs ON s.id = gs.shelf_id
		JOIN goods g ON gs.good_id = g.id
		JOIN goods_orders go ON g.id = go.good_id
		JOIN orders o ON go.order_id = o.id
		LEFT JOIN goods_shelves gs_additional ON gs.good_id = gs_additional.good_id AND gs_additional.main_or_additional = 'дополнительный'
		LEFT JOIN shelves sh ON gs_additional.shelf_id = sh.id
		WHERE
	`

	for i := 0; i < len(orderNumbers); i++ {
		if i == 0 {
			query += "o.number = " + orderNumbers[i]
		} else {
			query += " or o.number = " + orderNumbers[i]
		}
	}

	query += fmt.Sprint(" GROUP BY s.name, g.id, g.name, o.number, go.good_count;")

	rows, err := g.pool.Query(context.Background(), query)
	if err != nil {
		log.Println("Query")
		return nil, err
	}
	defer rows.Close()

	result := make(dto.AllGoodsByOrders)
	for rows.Next() {
		var shelfName string
		var good dto.GoodWithOrders
		var additionalShelf sql.NullString

		if err = rows.Scan(&shelfName, &good.Good.ID, &good.Good.Name, &good.OrderNumber, &good.GoodsCount, &additionalShelf); err != nil {
			return nil, err
		}

		if additionalShelf.Valid {
			good.AdditionalShelves = additionalShelf.String
		}

		result[shelfName] = append(result[shelfName], good)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func NewGood(pool *pgxpool.Pool) *GoodRepoImpl {
	return &GoodRepoImpl{pool: pool}
}
