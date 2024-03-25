package service

import (
	"github.com/khussa1n/shop/internal/dto"
	"github.com/khussa1n/shop/internal/model"
	"github.com/khussa1n/shop/internal/repository"
)

type GoodServiceImpl struct {
	repos *repository.Repository
}

func (g *GoodServiceImpl) GetAllByOrders(orderNumbers ...string) (dto.AllGoodsByOrders, error) {
	allGoodsByOrders := make(dto.AllGoodsByOrders)

	orders, err := g.repos.Order.GetAllByIds(orderNumbers...)
	if err != nil {
		return nil, err
	}

	orderIds := make([]int64, len(orders))
	for i, order := range orders {
		orderIds[i] = order.ID
	}

	goods_orders, err := g.repos.GoodOrder.GetAllByOrderIds(orderIds...)
	if err != nil {
		return nil, err
	}

	uniqueGoodsIds := make([]int64, len(goods_orders))
	uniqueGoodsIDs_map := make(map[int64]bool)

	for _, item := range goods_orders {
		if _, ok := uniqueGoodsIDs_map[item.GoodID]; !ok {
			uniqueGoodsIds = append(uniqueGoodsIds, item.GoodID)
			uniqueGoodsIDs_map[item.GoodID] = true
		}
	}

	goods, err := g.repos.Good.GetAllByIds(uniqueGoodsIds...)
	goods_map := make(map[int64]model.Goods)

	for _, item := range goods_orders {
		for _, good := range goods {
			if item.GoodID == good.ID {
				goods_map[item.GoodID] = good

			}
		}
	}

	goods_shelves, err := g.repos.GoodShelf.GetAllByGoodIds(uniqueGoodsIds...)
	if err != nil {
		return nil, err
	}

	main_shelvesIds := make([]int64, 0)
	additional_shelvesIds := make([]int64, 0)
	goods_additional_shelvesIds := make(map[int64][]int64)

	uniqueMainShelvesIDs := make(map[int64]bool)
	uniqueAdditionalShelvesIDs := make(map[int64]bool)

	for _, item := range goods_shelves {
		if item.MainOrAdditional == "главный" && !uniqueMainShelvesIDs[item.ShelfID] {
			main_shelvesIds = append(main_shelvesIds, item.ShelfID)
			uniqueMainShelvesIDs[item.ShelfID] = true
		}
		if item.MainOrAdditional != "главный" && !uniqueAdditionalShelvesIDs[item.ShelfID] {
			additional_shelvesIds = append(additional_shelvesIds, item.ShelfID)
			uniqueAdditionalShelvesIDs[item.ShelfID] = true
			goods_additional_shelvesIds[item.GoodID] = append(goods_additional_shelvesIds[item.GoodID], item.ShelfID)
		}
	}

	main_shelves_map := make(map[int64][]int64)
	for _, item1 := range goods_orders {
		for _, item2 := range goods_shelves {
			if item1.GoodID == item2.GoodID && item2.MainOrAdditional == "главный" {
				main_shelves_map[item2.ShelfID] = append(main_shelves_map[item2.ShelfID], item1.GoodID)
			}
		}
	}

	main_shelves, err := g.repos.Shelf.GetAllByIds(main_shelvesIds...)
	if err != nil {
		return nil, err
	}

	main_shelves_map2 := make(map[int64]string)
	for _, item := range main_shelves {
		main_shelves_map2[item.ID] = item.Name
	}

	additional_shelves, err := g.repos.Shelf.GetAllByIds(additional_shelvesIds...)
	if err != nil {
		return nil, err
	}

	goods_additional_shelves := make(map[int64][]string)
	for key, shelvesIds := range goods_additional_shelvesIds {
		for _, id := range shelvesIds {
			for _, shelf := range additional_shelves {
				if id == shelf.ID {
					goods_additional_shelves[key] = append(goods_additional_shelves[key], shelf.Name)
					break
				}
			}
		}
	}

	orderMap := make(map[int64]int64)
	for _, order := range orders {
		orderMap[order.ID] = order.Number
	}

	displayedGoods := make(map[int64]map[int64]bool)

	for key, good_ids := range main_shelves_map {
		for _, good_id := range good_ids {
			if _, ok := displayedGoods[good_id][key]; !ok {
				if displayedGoods[good_id] == nil {
					displayedGoods[good_id] = make(map[int64]bool)
				}

				displayedGoods[good_id][key] = true

				var orderNumbers []int64
				var goodsCounts []int64
				for _, item := range goods_orders {
					if item.GoodID == good_id {
						orderNumbers = append(orderNumbers, orderMap[item.OrderID])
						goodsCounts = append(goodsCounts, item.GoodCount)
					}
				}

				for i := range orderNumbers {
					allGoodsByOrders[main_shelves_map2[key]] = append(allGoodsByOrders[main_shelves_map2[key]],
						dto.GoodWithOrders{
							Good:              goods_map[good_id],
							OrderNumber:       orderNumbers[i],
							GoodsCount:        goodsCounts[i],
							AdditionalShelves: goods_additional_shelves[good_id],
						})
				}
			}
		}
	}

	return allGoodsByOrders, nil
}

func NewGood(repos *repository.Repository) *GoodServiceImpl {
	return &GoodServiceImpl{repos: repos}
}
