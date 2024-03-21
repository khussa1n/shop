package service

import (
	"github.com/khussa1n/shop/internal/dto"
	"github.com/khussa1n/shop/internal/repository"
)

type GoodServiceImpl struct {
	repos *repository.Repository
}

func (g *GoodServiceImpl) GetAllByOrders(orderNumbers ...string) (dto.AllGoodsByOrders, error) {
	return g.repos.Good.GetAllByOrders(orderNumbers...)
}

func NewGood(repos *repository.Repository) *GoodServiceImpl {
	return &GoodServiceImpl{repos: repos}
}
