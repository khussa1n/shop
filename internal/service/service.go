package service

import (
	"github.com/khussa1n/shop/internal/dto"
	"github.com/khussa1n/shop/internal/repository"
)

type Service struct {
	Good GoodServiceI
}

type GoodServiceI interface {
	GetAllByOrders(orderNumbers ...string) (dto.AllGoodsByOrders, error)
}

func NewService(Repos *repository.Repository) *Service {
	goodS := NewGood(Repos)
	return &Service{
		Good: goodS,
	}
}
