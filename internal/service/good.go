package service

import (
	"github.com/khussa1n/shop/internal/dto"
	"github.com/khussa1n/shop/internal/repository"
)

type GoodServiceImpl struct {
	repos *repository.Repository
}

func (g *GoodServiceImpl) GetAllByOrders(orderNumbers ...string) (dto.AllGoodsByOrders, error) {
	ordersByGoods, err := g.repos.Good.GetOrdersByNumbers(orderNumbers...)
	if err != nil {
		return nil, err
	}

	var goodIDs []int64
	for _, order := range ordersByGoods {
		goodIDs = append(goodIDs, order.GoodID)
	}

	goods, err := g.repos.Good.GetAllGoodsByIds(goodIDs...)
	if err != nil {
		return nil, err
	}

	shelvesByGoods, err := g.repos.Good.GetShelvesByGoods(goodIDs...)
	if err != nil {
		return nil, err
	}

	result := make(dto.AllGoodsByOrders)
	for _, order := range ordersByGoods {
		for _, good := range goods {
			if good.ID == order.GoodID {
				var additionalShelves []string
				for _, shelves := range shelvesByGoods {
					if shelves.GoodID == good.ID {
						additionalShelves = shelves.AdditionalShelves
						break
					}
				}
				var key string
				for _, shelves := range shelvesByGoods {
					if shelves.GoodID == good.ID {
						key = shelves.MainShelf
						break
					}
				}
				result[key] = append(result[key], dto.GoodWithOrders{
					Good:              good,
					OrderNumber:       order.OrderNumber,
					GoodsCount:        order.GoodCount,
					AdditionalShelves: additionalShelves,
				})
				break
			}
		}
	}

	return result, nil
}

func NewGood(repos *repository.Repository) *GoodServiceImpl {
	return &GoodServiceImpl{repos: repos}
}
