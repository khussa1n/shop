package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/khussa1n/shop/internal/model"
)

type Repository struct {
	Good UserRepoI
}

type UserRepoI interface {
	GetAllGoodsByIds(ids ...int64) ([]model.Goods, error)
	GetOrdersByNumbers(numbers ...string) ([]model.OrdersByGoods, error)
	GetShelvesByGoods(ids ...int64) ([]model.ShelvesByGoods, error)
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	goodR := NewGood(pool)
	return &Repository{
		Good: goodR,
	}
}
