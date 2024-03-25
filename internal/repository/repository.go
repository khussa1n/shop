package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/khussa1n/shop/internal/model"
)

type Repository struct {
	Good      GoodRepoI
	Order     OrderRepoI
	Shelf     ShelfRepoI
	GoodOrder GoodOrderRepoI
	GoodShelf GoodShelfRepoI
}

type GoodRepoI interface {
	GetAllByIds(ids ...int64) ([]model.Goods, error)
}

type OrderRepoI interface {
	GetAllByIds(ids ...string) ([]model.Orders, error)
}

type ShelfRepoI interface {
	GetAllByIds(ids ...int64) ([]model.Shelves, error)
}

type GoodOrderRepoI interface {
	GetAllByOrderIds(ids ...int64) ([]model.GoodsOrders, error)
}

type GoodShelfRepoI interface {
	GetAllByGoodIds(ids ...int64) ([]model.GoodsShelves, error)
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	goodR := NewGood(pool)
	orderR := NewOrder(pool)
	shelfR := NewShelf(pool)
	goodOrderR := NewGoodOrder(pool)
	goodShelf := NewGoodShelf(pool)
	return &Repository{
		Good:      goodR,
		Order:     orderR,
		Shelf:     shelfR,
		GoodOrder: goodOrderR,
		GoodShelf: goodShelf,
	}
}
