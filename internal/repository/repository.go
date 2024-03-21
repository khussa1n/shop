package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/khussa1n/shop/internal/dto"
)

type Repository struct {
	Good UserRepoI
}

type UserRepoI interface {
	GetAllByOrders(orderNumbers ...string) (dto.AllGoodsByOrders, error)
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	goodR := NewGood(pool)
	return &Repository{
		Good: goodR,
	}
}
